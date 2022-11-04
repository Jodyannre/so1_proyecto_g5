import React, { useState,useEffect } from "react";
import axios from "axios";
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

/*
const UserData = [
  {
      id: 1,
      year: 2016,
      userGain: 80000,
      userLost: 823,
  },
  {
      id: 2,
      year: 2017,
      userGain: 45677,
      userLost: 345,
  },
  {
      id: 3,
      year: 2018,
      userGain: 78888,
      userLost: 555,
  },
  {
      id: 4,
      year: 2019,
      userGain: 90000,
      userLost: 4555,
  },
  {
      id: 5,
      year: 2020,
      userGain: 4300,
      userLost: 234,
  },
]
*/  

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


const labels = ['2-2', '3-2', '5-1', '0-0', '3-3', '1-2', '0-3'];
const numbers = [100,200,50,25,75,65,72]

const data = {
  labels,
  datasets: [
    {
      label: 'Resultados en vivo',
      data: numbers.map((a) => a),
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



export default function GraficaLive({match}) {
  const [datos, setDatos] = useState([])
  const [partido,setPartido] = useState(match)

  const getDataPartidos = async () => {
    const res = await axios.get(`http://${global.ip}:${global.port}/getDataPartidos`,{params: { key: partido}});
    //const texto = JSON.stringify(res.data)
    setDatos(res.data)
    //console.log(res.data)
    //console.log(datos)
  }

  useEffect(()=>{
    console.log("Si vino a ver partidos")
    console.log(match)
    if (match){
      setPartido(match)
    }
    const interval = setInterval(()=>{
      getDataPartidos();
    }, 5000)
    getDataPartidos();
  },[match]);


  return (
    <>  
      <div className="square bg-primary rounded-pill bg-light" style=
      {{hposition:"relative",marginBottom:"1%",padding:"1%",
        backgroundColor :"white"}}>
        <h1><center>Gráficas</center></h1>  
        </div>          
        <div style={{height:"60vh",position:"relative",marginBottom:"1%",padding:"1%",
        backgroundColor :"white", border:"4px dotted blue",
        overflow: "scroll"}}>
          <Bar options={options} data={data} />
        </div>
        </>

  )
}
