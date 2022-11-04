import React, { useState,useEffect } from "react";
import axios from "axios";
import '../styles/styles.css'
import Table from 'react-bootstrap/Table'
import Button from 'react-bootstrap/Button';
//import {Form} from 'react-bootstrap'
import Card from 'react-bootstrap/Card';
import Spinner from 'react-bootstrap/Spinner';

export default function MongoDash() {
  const [logs, setLogs] = useState([])
  const [total,setTotal] = useState(0)
  const [time, setTime] = React.useState(0);
  const [showPage, setShowPage] = useState(false)
  const [purgeOn, setPurgeOn] = useState(false)

  const getLogsMongo = async () => {
    const res = await axios.get(`http://${global.ip}:${global.port}/getLogsMongo`);
    //const texto = JSON.stringify(res.data)
    setLogs(res.data)
    //console.log(res.data)
    //console.log(texto)
    //console.log(datos)
    setShowPage(true)
  }

  const getTotalMongo = async () => {
    const res = await axios.get(`http://${global.ip}:${global.port}/getTotalMongo`);
    //const texto = JSON.stringify(res.data)
    setTotal(res.data.Total)
    console.log(res.data.Total)
    //console.log(datos)
  }

  const goPurge = async (e) =>{
    const res = await axios.get(`http://${global.ip}:${global.port}/purge`);
    setPurgeOn(!purgeOn)
    //console.log(e.target);
    //alert("Click en el boton")
    console.log(res.data)
  }

  const GetLogsFormat = ()=>{
    return (
        logs.map((value,key)=>{
            return (
                <tr>
                    <td>
                        <Card body>{JSON.stringify(value)}</Card>
                    </td>
                </tr> 
            ) 
        })
    )
}

  useEffect(()=>{

    getTotalMongo();
    getLogsMongo();

    const timer = window.setInterval(() => {
      setTime(time + 1);
    }, 1250);
    return () => {
      window.clearInterval(timer);
    };
  },[time,purgeOn]);

  //Tabla de resultados
  //const data= [{memoria:"Joel Rodriguez", pid:25,nombre:"Estudiante",usuario:"otro val"}]
  //const dataChildren= [{nombre:"no tiene",pid:-1}]
  //const headers = [{name:"nombre",age:"edad",profession:"otro", children:"otro"}]
  return (
        showPage ? <>
            <div className="square bg-primary rounded-pill bg-light" 
            style={{hposition:"relative",marginBottom:"1%",padding:"1%",
            backgroundColor :"white"}}>
            <center><h1>Últimos 10 Logs</h1></center>  
            </div>     
        <div className="MongoLog" style={{height:"350px",overflow: "scroll"}}>  
            <Table striped bordered hover >
                <thead>
                <tr>
                    <th>Datos</th>
                </tr>  
                </thead>
                <tbody>
                    {GetLogsFormat()}
                </tbody>              
            </Table> 
        </div>

        <div className="App" style={{height:"220px"}}>  
        <Table striped bordered hover size="sm" >
                <thead>
                <tr style={{textAlign:"center"}}>
                    <th style={{verticalAlign: "middle"}}>Total Records</th>
                    <th style={{verticalAlign: "middle"}}>Purge</th>
                </tr>  
                </thead>
                <tbody>
                        <tr>
                            <td style={{fontSize:"40px"}}>
                                {total}
                            </td>
                            <td style={{verticalAlign: "middle"}}>
                            <Button onClick={goPurge} variant="outline-primary">Purge</Button>{' '}
                            </td>
                        </tr>                    
                </tbody>              
            </Table>
        </div>
        </> : 
        <div style={{textAlign:"center", verticalAlign:"middle"}}>
            <Spinner animation="border" role="status" variant="light">
            <span className="visually-hidden">Loading...</span>
            </Spinner>            
        </div>
  )
}
