import {useState} from 'react';
import logo from './assets/images/logo.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import {Grid} from "@mui/material";

function App() {
  const [resultText, setResultText] = useState("Dump all YAMLs 👇");
  const [name, setName] = useState('');
  const updateName = (e: any) => setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);

  function greet() {
    Greet(name).then(updateResultText);
  }


  return (
    <div id="App">
      <Grid container spacing={2}>
        <Grid item xs={4}>
          <img src={logo} id="logo" alt="logo"/>
        </Grid>
        <Grid item xs={8}>
          <div id="input">
            <button className="btn" onClick={greet}>DUMP Y.A.M.L</button>
          </div>
        </Grid>
        <Grid item xs={12}>
          <p>
            {resultText}
          </p>
        </Grid>
      </Grid>

    </div>
  )
}

export default App
