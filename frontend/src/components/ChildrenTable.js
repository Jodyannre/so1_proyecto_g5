import React, { useState,useEffect } from "react";
import '../styles/styles.css'
import Table from 'react-bootstrap/Table'
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";


export default function ChildrenTable({datos}) {
  const [data, setData] = useState(datos)
  const [show, setShow] = useState(false);
  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  useEffect(()=>{
    const interval = setInterval(()=>{
    }, 20000)
    setData(datos)
    setShow(false)
  },[]);

  //Tabla de resultados
  return (
        <div className="App" >  
<Button variant="primary" onClick={handleShow}>
        Mostrar
      </Button>
      <Modal show={show} onHide={handleClose} scrollable={true}>
        <Modal.Header closeButton>
          <Modal.Title>Procesos hijos</Modal.Title>
        </Modal.Header>
        <Modal.Body>
        <Table striped bordered hover >
                <thead>
                <tr>
                    <th>PID</th>
                    <th>Nombre</th>
                </tr>  
                </thead>
                <tbody>
                    {data.map((val,key)=>{
                        return(
                        <tr key={key}>
                            <td>{val.pid}</td>
                            <td>{val.nombre}</td>
                        </tr>
                        )
                    })}
                </tbody>              
            </Table>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Cerrar
          </Button>
        </Modal.Footer>
      </Modal>
        </div>
  )
}
