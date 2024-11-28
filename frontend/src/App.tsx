import {useState} from 'react';
import logo from './assets/images/logo.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";

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
      <img src={logo} id="logo" alt="logo"/>
      {/*<div id="result" className="result">{resultText}</div>*/}
      <div id="input">
        <button className="btn" onClick={greet}>DUMP Y.A.M.L</button>
      </div>
      <p>
        {resultText}
      </p>
    </div>
  )
}

export default App
