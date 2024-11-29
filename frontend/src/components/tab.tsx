import * as React from 'react';
import {useEffect} from 'react';
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';
import Box from '@mui/material/Box';
import {pkg} from '../../wailsjs/go/models';
import {AppBar, Typography} from "@mui/material";
import NestedList from "./list";

interface TabPanelProps {
  profile: pkg.Profile;
}

function TabPanel(props: { children?: React.ReactNode; value: string; index: string }) {
  const {children, value, index, ...other} = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`tabpanel-${index}`}
      aria-labelledby={`tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box sx={{p: 3}}>
          <Typography>{children}</Typography>
        </Box>
      )}
    </div>
  );
}

export default function MyTabs(props: TabPanelProps) {
  const [value, setValue] = React.useState("one");
  const keys = Object.keys(props.profile);
  const [apLightsProfile, setApLightsProfile] = React.useState({} as pkg.Profile);
  const [apKnobsProfile, setApKnobsProfile] = React.useState({} as pkg.Profile);


  const handleChange = (event: React.SyntheticEvent, newValue: string) => {

    let apLightsProfileCopy: any = {
      metadata: props.profile.metadata,
      ap: props.profile.ap,
      hdg: props.profile.hdg,
      nav: props.profile.nav,
      apr: props.profile.apr,
      rev: props.profile.rev,
      alt: props.profile.alt,
      vs: props.profile.vs,
      ias: props.profile.ias,
    }
    let apKnobsProfileCopy: any = {
      metadata: props.profile.metadata,
      ap_hdg: props.profile.ap_hdg,
      ap_vs: props.profile.ap_vs,
      ap_alt: props.profile.ap_alt,
      ap_ias: props.profile.ap_ias,
      ap_crs: props.profile.ap_crs,

    }
    setApLightsProfile(apLightsProfileCopy);
    setApKnobsProfile(apKnobsProfileCopy);
    setValue(newValue);
    console.log(newValue)
  };

  useEffect(() => {
    handleChange({} as React.SyntheticEvent, "one");
  }, [props.profile]);

  return (
    <Box sx={{width: '100%', overflow: "auto"}}>
      <AppBar position="static">
        <Tabs
          value={value}
          onChange={handleChange}
          textColor="inherit"
        >
          <Tab value="one" label="AP Lights"/>
          <Tab value="two" label="Lights 1"/>
          <Tab value="three" label="Lights 2"/>
          <Tab value="four" label="AP Knobs"/>
          <Tab value="five" label="Button*" disabled={true}/>
        </Tabs>
      </AppBar>
      <Box>
        <TabPanel value={value.toString()} index="one">
          <NestedList profile={apLightsProfile}/>
        </TabPanel>
        <TabPanel value={value.toString()} index="two">
          {keys.map((key, index) => {
            if (key !== "metadata") {
              // @ts-ignore
              if (props.profile[key].profile_type === "knob") {


                return (
                  <>
                    <Typography key={index}>{key}-{
                      // @ts-ignore
                      JSON.stringify(props.profile[key])
                    }</Typography>
                  </>
                )
              }
            }
          })}
        </TabPanel>
        <TabPanel value={value.toString()} index="three">

        </TabPanel>
        <TabPanel value={value.toString()} index="four">
          <NestedList profile={apKnobsProfile}/>
        </TabPanel>
      </Box>
    </Box>
  );
}