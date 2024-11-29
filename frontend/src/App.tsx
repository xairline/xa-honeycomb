import React, {useEffect, useState} from 'react';
import './App.css';
import {GetProfile, GetProfiles} from "../wailsjs/go/main/App";
import {Grid, Stack} from "@mui/material";
import {pkg} from "../wailsjs/go/models";
import Profiles from "./components/profiles";
import Metadata from "./components/metadata";
import Lights from './components/lights';

function App() {
  const [profileData, setProfileData] = useState({} as pkg.Profile);
  const [profilesData, setProfilesData] = useState([] as pkg.Profile[]);
  const [name, setName] = useState('');
  const profile = (result: pkg.Profile) => setProfileData(result);
  const profiles = (result: pkg.Profile[]) => setProfilesData(result);
  const [selectedIndex, setSelectedIndex] = React.useState(0);

  const handleListItemClick = (
    event: React.MouseEvent<HTMLDivElement, MouseEvent>,
    index: number,
  ) => {
    setSelectedIndex(index);
    GetProfile(profilesData[index].metadata?.name || "").then(profile);
  };

  useEffect(() => {
    GetProfiles().then(profiles);
  }, []);

  useEffect(() => {
    if (profilesData.length > 0) {
      GetProfile(profilesData[selectedIndex].metadata?.name || "").then(profile);
    }
  }, [profilesData]);


  return (
    <div id="App">
      <Grid container spacing={0} columns={16} style={{height: "100vh", overflow: "hidden"}}>
        <Grid item xs={4} style={{border: "white", overflow: "hidden"}}>
          <Profiles profiles={profilesData} selectedIndex={selectedIndex} handleListItemClick={handleListItemClick}/>
        </Grid>
        <Grid item xs={12} sx={{width: '100vw', overflow: "hidden", height: "100vh"}}>
          <Grid item xs={16}
                sx={{width: '100vw', overflow: "auto", bgcolor: "#222e35", height: "100vh"}}>
            <Stack spacing={2} sx={{margin: "18px"}}>
              <Metadata metadata={profileData?.metadata}/>
              <Lights lights={profileData.leds} title={"Auto Pilot Lights"}/>
              <Lights title={"Annunciators Row (Top)"}/>
              <Lights title={"Annunciators Row (Bottom)"}/>
              <Lights title={"Auto Pilot Knobs"}/>
              <Lights title={"Landing Gear Lights"}/>
            </Stack>

          </Grid>
        </Grid>
      </Grid>

    </div>
  )
}

export default App
