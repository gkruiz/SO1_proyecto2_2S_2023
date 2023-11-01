

import React, { useEffect , useState} from "react";
import { Link ,useNavigate } from "react-router-dom";
import './App.css';
import { CircularProgressbar ,buildStyles} from 'react-circular-progressbar';
import 'react-circular-progressbar/dist/styles.css';
import io from "socket.io-client";





/*INICIA PARTE GRAFICA EN EL PROGRAMA */

 function App() {
  //const [products, setProducts] = useState(dataResponse);
  const [mensaje, setMensaje] = useState([]);


  const ws = "ws://127.0.0.1:3000/";

  useEffect(() => {
    // connect to WebSocket server
    const newSocket = io(ws);
    // set up event listeners for incoming messages
    newSocket.on("connect", () => console.log("Connected to WebSocket"));
    newSocket.on("disconnect", () =>
      console.log("Disconnected from WebSocket")
    );
    newSocket.on("message", (data) => {

      setMensaje(JSON.parse(data));


    });

    // clean up on unmount
    return () => {
      newSocket.disconnect();
    };
  });





    

console.log(mensaje);
   

  return (
   
      
    <div className="principal" >
    <table class="table table-dark table-striped"  >
      <thead>
        <tr>
            <td colSpan={6}>
              <h4>Muestra datos activos Redis</h4>
            </td>
        </tr>
      </thead>
      <tbody>
        {
          mensaje.map((dat, index)=>{
            return(
              <tr key={index}>
                <td>
                    {dat.album}
                </td>
                <td>
                    {dat.artist}
                </td>
                <td>
                    {dat.year}
                </td>
              </tr>
            )

          })

        }
        
        


      </tbody>



    </table>





  </div>
    


  );
}

export default App;










