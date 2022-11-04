import React, { useState,useEffect } from "react";
import axios from "axios";
import '../styles/styles.css'
import Table from 'react-bootstrap/Table'
import {Form} from 'react-bootstrap'
//import Dropdown from 'react-bootstrap/Dropdown';
//import GraficaLive from "./GraficaLive";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';
import { Bar } from 'react-chartjs-2';


ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
);

  
const options = {
  indexAxis: 'y',
  elements: {
    bar: {
      borderWidth: 2,
    },
  },
  responsive: true,
  plugins: {
    legend: {
      position: 'right',
    }/*,
    title: {
      display: true,
      text: 'Resultado registrados',
    },*/
  },
}


//const labels = ["2-2", "3-2", "5-1"];
//const numbers = [100,200,50]


export default function SelectorPartido() {
  const [datos, setDatos] = useState([])
  const [fases, setFases] = useState([])
  const [partido,setPartido] = useState("Seleccione un partido")
  const [key, setKey] = useState("")
  //const [prekey, setPrekey] = useState("")
  const [selectedPhase, setSelectedPhase] = useState(0)
  //const [llamarGrafica, setLlamarGrafica] = useState(false)
  const [time, setTime] = React.useState(0);
  //const [dataGraph, setDataGraph] = useState([])
  //const [labelsGraph,setLabelsGraph] = useState([])
  //const [tituloGraph, setTituloGraph] = useState("")


  const getDataGraph = () => {
    var dato = [0.5]
    var label = [""]
    label.pop()
    dato.pop()
    try{
    datos.map((obj) => {
      //console.log(obj)
      dato.push(obj.Total)
      label.push(obj.Partido)
      return 1
    })
    }catch(err){
      //console.log("Medio error, porque no se selecciono partido, en fin, no es error.")
    }
    //console.log(dato)
    //console.log(label)
    //console.log("Si entro aqui")
    dato.push(100)
    const data = {
      labels:label,
      datasets: [
        {
          label: 'live',
          data: dato.map((a) => a),
          borderColor: [
            'rgba(255, 99, 132, 1)',
            'rgba(54, 162, 235, 1)',
            'rgba(255, 206, 86, 1)',
            'rgba(75, 192, 192, 1)',
            'rgba(153, 102, 255, 1)',
            'rgba(255, 159, 64, 1)',
          ],
          backgroundColor: [
            'rgba(255, 99, 132, 0.2)',
            'rgba(54, 162, 235, 0.2)',
            'rgba(255, 206, 86, 0.2)',
            'rgba(75, 192, 192, 0.2)',
            'rgba(153, 102, 255, 0.2)',
            'rgba(255, 159, 64, 0.2)',
          ]
        }/*,
        {
          label: 'Dataset 2',
          data: numbers.map((a) => a),
          borderColor: 'rgb(53, 162, 235)',
          backgroundColor: 'rgba(53, 162, 235, 0.5)',
        },*/
      ],
    };
    return data
  }




  const getDataFases = async () => {
    const res = await axios.get(`http://${global.ip}:${global.port}/getDataFases`);
    //const texto = JSON.stringify(res.data)
    res.data.Fase1.sort()
    res.data.Fase2.sort()
    res.data.Fase3.sort()
    res.data.Fase4.sort()
    setFases(res.data)
    //const array = res.data
    console.log(fases)
    //console.log(res.data.Fase2)
  }

  const getDataPartidos = async () => {
    const res = await axios.get(`http://${global.ip}:${global.port}/getDataPartidos`,{params: { key: key}});
    //const texto = JSON.stringify(res.data)
    setDatos(res.data)
    console.log(res.data)
    //console.log(datos)
  }
  const getPartidosFase = () =>{
    switch(selectedPhase){
      case "1":
        return(
          fases.Fase1 ? 
            fases.Fase1.map((val,key)=>{
              return (<option value={val}>{val}</option>)
            }) : null 
            )
      case "2":
        return(
          fases.Fase2 ? 
            fases.Fase2.map((val,key)=>{
              return (<option value={val}>{val}</option>)
            }) : null 
            )
      case "3":
        return(
          fases.Fase3 ? 
            fases.Fase3.map((val,key)=>{
              return (<option value={val}>{val}</option>)
            }) : null 
            )
      case "4":
        return(
          fases.Fase4 ? 
            fases.Fase4.map((val,key)=>{
              return (<option value={val}>{val}</option>)
            }) : null 
           )
        default:
          return null
    }
}

const onChangePhase = (e)=>{
  document.getElementById("selector-partido").selectedIndex = 0
  setSelectedPhase(e.target.value);
  setKey("")
  setDatos([])
  setPartido("Seleccione un partido")
}

const onChangeKey = (e)=>{
  setKey(e.target.value+":"+selectedPhase)
  var str = ""
  str = e.target.value.toString()
  setPartido("Gráfica de "+ str.replace(":"," vs "))
}

  useEffect(()=>{
    console.log(key)
    if (key!==""){
      getDataPartidos();
    }
    getDataFases();
    const timer = window.setInterval(() => {
      setTime(time + 1);
      /*
      if (key!=""){
        console.log("Key dentro del interval")
        console.log(key)
        getDataPartidos();
        getDataFases();
      }
      */
    }, 3500);
    return () => {
      window.clearInterval(timer);
    };
  },[key,time]);

  return (
        <>
            <div className="square bg-primary rounded-pill bg-light" style={{hposition:"relative",marginBottom:"1%",padding:"1%",
            backgroundColor :"white"}}>
            <center><h1>Live</h1>
              <lord-icon
                        src="https://cdn.lordicon.com/fbyjqbak.json"
                        trigger="loop"
                        colors="primary:#e83a30"
                        style={{width:"50px",height:"50px"}}>
                    </lord-icon>
            </center>  
            
            </div>     
        <div className="App" style={{height:"210px"}}>  

            <Table striped bordered hover  >
                <thead>
                <tr>
                    <th >Fase</th>
                    <th >Partido</th>
                </tr>  
                </thead>
                <tbody>
                        <tr>
                            <td>
                            <Form.Select                       
                            onChange={(e)=>{onChangePhase(e)}}
                            >                            
                            <option>Seleccionar fase</option>
                            <option value={1}>1/16</option>
                            <option value={2}>Octavos de final</option>
                            <option value={3}>Semifinal</option>
                            <option value={4}>Final</option>
                            </Form.Select>
                            </td>
                            <td>
                            <Form.Select
                            id="selector-partido"
                            onChange={(e)=>{
                            onChangeKey(e)
                            }}
                            >
                              <option>Seleccionar partido</option>
                              {getPartidosFase()}
                            </Form.Select>
                            </td>
                        </tr>                    
                </tbody>              
            </Table>
        </div>
        <div className="square bg-primary rounded-pill bg-light" style=
      {{hposition:"relative",marginBottom:"1%",padding:"1%",
        backgroundColor :"white"}}>
        <h1><center>{partido}</center></h1>  
        </div>          
        <div style={{height:"85vh",position:"relative",marginBottom:"1%",padding:"1%",
        backgroundColor :"white", border:"4px dotted blue",
        overflow: "scroll"}}>
          <Bar options={options} data={getDataGraph()} />
        </div>
        </>
  )
}
