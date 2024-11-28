import {useState} from 'react';
import './App.css';
import {GetProfiles, Greet} from "../wailsjs/go/main/App";
import {Grid} from "@mui/material";

function App() {
  const [resultText, setResultText] = useState("Dump all YAMLs 👇");
  const [resultText2, setResultText2] = useState("Yamls");
  const [name, setName] = useState('');
  const updateName = (e: any) => setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);
  const updateResultText2 = (result: string) => setResultText2(JSON.stringify(result));

  function greet() {
    Greet(name).then(updateResultText);
    GetProfiles().then(updateResultText2);
  }


  return (
    <div id="App">
      <Grid container spacing={2} columns={16} style={{height: "100vh"}}>
        <Grid item xs={4} style={{border: "white"}}>
          {resultText2}
        </Grid>
        <Grid item xs={12}>
          <Grid item xs={16}>
            <div id="input">
              <button className="btn" onClick={greet}>DUMP Y.A.M.L</button>
            </div>
          </Grid>
          <Grid item xs={16}>
            <p>
              {resultText}
            </p>
          </Grid>
        </Grid>
      </Grid>

    </div>
  )
}

export default App
