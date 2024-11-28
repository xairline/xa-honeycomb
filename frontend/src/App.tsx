import React, {useEffect, useState} from 'react';
import './App.css';
import {GetProfile, GetProfiles} from "../wailsjs/go/main/App";
import {Box, Grid, List, ListItemButton, ListItemIcon, ListItemText} from "@mui/material";
import {pkg} from "../wailsjs/go/models";
import InboxIcon from '@mui/icons-material/Inbox';
import MyTabs from "./components/tab";

function App() {
  const [profileData, setProfileData] = useState({} as pkg.Profile);
  const [profilesData, setProfilesData] = useState([] as pkg.Profile[]);
  const [name, setName] = useState('');
  const profile = (result: pkg.Profile) => setProfileData(result);
  const profiles = (result: pkg.Profile[]) => setProfilesData(result);
  const [selectedIndex, setSelectedIndex] = React.useState(1);

  const handleListItemClick = (
    event: React.MouseEvent<HTMLDivElement, MouseEvent>,
    index: number,
  ) => {
    setSelectedIndex(index);
    GetProfile(profilesData[index].metadata.name || "").then(profile);
  };

  useEffect(() => {
    GetProfiles().then(profiles);
  }, []);


  return (
    <div id="App">
      <Grid container spacing={0} columns={16} style={{height: "100vh", overflow: "hidden"}}>
        <Grid item xs={4} style={{border: "white", overflow: "hidden"}}>
          <Box sx={{width: '100%', overflow: "hidden", bgcolor: "#022e35", height: "100vh"}}>
            < List component="nav" aria-label="main mailbox folders">
              {profilesData.map((profile, index) => {
                return (
                  <ListItemButton
                    selected={selectedIndex === index}
                    onClick={(event) => handleListItemClick(event, index)}
                  >
                    <ListItemIcon>
                      <InboxIcon/>
                    </ListItemIcon>
                    <ListItemText primary={profile.metadata.name}/>
                  </ListItemButton>
                )
              })}
            </List>

          </Box>
        </Grid>
        <Grid item xs={12} sx={{width: '100vw', overflow: "hidden", height: "100vh"}}>
          <Grid item xs={16}
                sx={{width: '100vw', overflow: "hidden", bgcolor: "#222e35", height: "80vh"}}>
            <MyTabs profile={profileData}/>
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
