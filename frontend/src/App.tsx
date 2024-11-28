import React, {useEffect, useState} from 'react';
import './App.css';
import {GetProfiles, Greet} from "../wailsjs/go/main/App";
import {Box, Grid, List, ListItemButton, ListItemIcon, ListItemText} from "@mui/material";
import {pkg} from "../wailsjs/go/models";
import InboxIcon from '@mui/icons-material/Inbox';

function App() {
  const [resultText, setResultText] = useState({} as pkg.Profile);
  const [resultText2, setResultText2] = useState([] as pkg.Profile[]);
  const [name, setName] = useState('');
  const updateResultText = (result: pkg.Profile) => setResultText(result);
  const updateResultText2 = (result: pkg.Profile[]) => setResultText2(result);
  const [selectedIndex, setSelectedIndex] = React.useState(1);

  const handleListItemClick = (
    event: React.MouseEvent<HTMLDivElement, MouseEvent>,
    index: number,
  ) => {
    setSelectedIndex(index);
    console.log(resultText2[index].name);
    Greet(resultText2[index].name).then(updateResultText);
  };

  function greet() {
    Greet(name).then(updateResultText);
    GetProfiles().then(updateResultText2);
  }

  useEffect(() => {
    GetProfiles().then(updateResultText2);
  }, []);


  return (
    <div id="App">
      <Grid container spacing={0} columns={16} style={{height: "100vh", overflow: "hidden"}}>
        <Grid item xs={4} style={{border: "white", overflow: "hidden"}}>
          <Box sx={{width: '100%', overflow: "hidden", bgcolor: "#022e35", height: "100vh"}}>
            < List component="nav" aria-label="main mailbox folders">
              {resultText2.map((profile, index) => {
                return (
                  <ListItemButton
                    selected={selectedIndex === index}
                    onClick={(event) => handleListItemClick(event, index)}
                  >
                    <ListItemIcon>
                      <InboxIcon/>
                    </ListItemIcon>
                    <ListItemText primary={profile.name}/>
                  </ListItemButton>
                )
              })}
            </List>

          </Box>
        </Grid>
        <Grid item xs={12} sx={{width: '100vw', overflow: "hidden", height: "100vh"}}>
          <Grid item xs={16}
                sx={{width: '100vw', overflow: "hidden", bgcolor: "#222e35", height: "80vh"}}>
            <p>
              {JSON.stringify(resultText)}
            </p>
          </Grid>
          <Grid item xs={16}
                sx={{width: '100vw', overflow: "hidden", bgcolor: "#122e35", height: "20vh"}}>
            <Box>
              <p>TODO: XPLANE STATUS AND LOGS</p>
            </Box>
          </Grid>
        </Grid>
      </Grid>

    </div>
  )
}

export default App
