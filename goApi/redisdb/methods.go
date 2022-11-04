package redisdb

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strconv"

	"github.com/go-redis/redis/v9"
)

/*
	func GetRedisInfo() {
		var ctx = context.Background()
		ip := hostRedis + ":" + portRedis
		rdb := redis.NewClient(&redis.Options{
			Addr:     ip,
			Password: "",
			DB:       0,
		})
		//Obtener datos
		var key = fmt.Sprintf("%s:%s:%s",t1,t2,phase)
		var equipos = fmt.Sprintf("%s:%s",t1,t2)
		var fase = fmt.Sprintf("Fase:%s",phase)
		var keyTotal = fmt.Sprintf("%s:%s:%s:total",t1,t2,phase)


		//Agregar o incrementar  valor
		val := rdb.HIncrBy(ctx, key, field, 1)
		total := rdb.IncrBy(ctx,keyTotal,1)
		conjunto := rdb.SAdd(ctx,fase,equipos)
	}
*/
func SaveMatch(client *redis.Client, ctx context.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		SetCors(res)
		//Obtener info
		var match Match
		var ctx = context.Background()
		json.NewDecoder(req.Body).Decode(&match)

		//Obtener datos
		var key = fmt.Sprintf("%s:%s:%d", match.Team1, match.Team2, match.Phase)
		var equipos = fmt.Sprintf("%s:%s", match.Team1, match.Team2)
		var fase = fmt.Sprintf("Fase:%d", match.Phase)
		var keyTotal = fmt.Sprintf("%s:%s:%d:total", match.Team1, match.Team2, match.Phase)
		var field = match.Score

		//Agregar o incrementar  valor
		val := client.HIncrBy(ctx, key, field, 1)
		total := client.IncrBy(ctx, keyTotal, 1)
		conjunto := client.SAdd(ctx, fase, equipos)
		fmt.Println("El valor guardado  es: ", val)
		fmt.Println("El total es: ", total)
		fmt.Println("El estado del conjunto es: ", conjunto)
		json.NewEncoder(res).Encode("Se registró el valor correctamente.")
	}
}

func GetAllMatches(client *redis.Client, ctx context.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		SetCors(res)
		//Obtener info
		var fases Fases

		//Guardar fases
		fases.Fase1 = GetFase("Fase:1", client, ctx)
		fases.Fase2 = GetFase("Fase:2", client, ctx)
		fases.Fase3 = GetFase("Fase:3", client, ctx)
		fases.Fase4 = GetFase("Fase:4", client, ctx)
		json.NewEncoder(res).Encode(fases)
	}
}

func GetCounters(client *redis.Client, ctx context.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		partido := req.URL.Query().Get("key")
		fmt.Println(partido)
		SetCors(res)
		//var key Key
		var resultados Resultados
		//var grafica Grafica
		//json.NewDecoder(req.Body).Decode(&key)
		result := client.HGetAll(ctx, partido).Val()
		//Conseguir total de partidos para calcular porcentaje
		result2 := client.Get(ctx, partido+":total")
		totalPartidos, _ := strconv.Atoi(result2.Val())
		//Crear un slide con las keys
		keys := make([]string, 0, len(result))
		for k := range result {
			keys = append(keys, k)
		}

		//Ordenar el array
		sort.Strings(keys)

		for i := 1; i <= len(keys); i++ {
			value := result[keys[len(keys)-i]]
			total, _ := strconv.ParseFloat(value, 64)
			nConteo := Conteo{Partido: keys[len(keys)-i], Total: math.Round((total / float64(totalPartidos)) * 100)}
			resultados = append(resultados, nConteo)
			//grafica.Data = append(grafica.Data, nConteo.Total)
			//grafica.Labels = append(grafica.Labels, nConteo.Partido)
		}
		//Orderan el resultado
		fmt.Println(resultados)
		json.NewEncoder(res).Encode(resultados)
	}
}

func GetFase(fase string, client *redis.Client, ctx context.Context) []string {
	return client.SMembers(ctx, fase).Val()
}

func SetCors(wri http.ResponseWriter) {
	(wri).Header().Set("Access-Control-Allow-Origin", "*")
	(wri).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(wri).Header().Set("Content-Type", "application/json")
}
