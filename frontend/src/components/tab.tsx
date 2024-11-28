import * as React from 'react';
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';
import Box from '@mui/material/Box';
import {pkg} from '../../wailsjs/go/models';
import {List, ListItemButton, ListItemIcon, ListSubheader, Typography} from "@mui/material";
import ListItemText from '@mui/material/ListItemText';
import InboxIcon from '@mui/icons-material/Inbox';
import DraftsIcon from '@mui/icons-material/Drafts';
import SendIcon from '@mui/icons-material/Send';
import ExpandLess from '@mui/icons-material/ExpandLess';
import ExpandMore from '@mui/icons-material/ExpandMore';
import StarBorder from '@mui/icons-material/StarBorder';
import Collapse from '@mui/material/Collapse';

interface TabPanelProps {
  profile: pkg.Profile;
  children?: React.ReactNode;
  value?: string;
  index?: string;
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
  const [value, setValue] = React.useState(0);
  const keys = Object.keys(props.profile);
  const [open, setOpen] = React.useState(true);

  const handleClick = () => {
    setOpen(!open);
  };
  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setValue(newValue);
  };

  return (
    <Box sx={{width: '100%'}}>
      <Tabs
        value={value}
        onChange={handleChange}
        textColor="secondary"
        indicatorColor="secondary"
        aria-label="secondary tabs example"
      >
        <Tab value="one" label="LED"/>
        <Tab value="two" label="Knob"/>
        <Tab value="three" label="Autopilot"/>
      </Tabs>
      <TabPanel value={value.toString()} index="one">
        {keys.map((key, index) => {
          if (key !== "metadata") {
            // @ts-ignore
            if (props.profile[key].profile_type === "led") {

              return (
                // @ts-ignore
                <Typography key={index}>{key}: {JSON.stringify(props.profile[key])}</Typography>
              )
            }
          }
        })}
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
                  <List
                    sx={{width: '100%', maxWidth: 360, bgcolor: 'background.paper'}}
                    component="nav"
                    aria-labelledby="nested-list-subheader"
                    subheader={
                      <ListSubheader component="div" id="nested-list-subheader">
                        Nested List Items
                      </ListSubheader>
                    }
                  >
                    <ListItemButton>
                      <ListItemIcon>
                        <SendIcon/>
                      </ListItemIcon>
                      <ListItemText primary="Sent mail"/>
                    </ListItemButton>
                    <ListItemButton>
                      <ListItemIcon>
                        <DraftsIcon/>
                      </ListItemIcon>
                      <ListItemText primary="Drafts"/>
                    </ListItemButton>
                    <ListItemButton onClick={handleClick}>
                      <ListItemIcon>
                        <InboxIcon/>
                      </ListItemIcon>
                      <ListItemText primary="Inbox"/>
                      {open ? <ExpandLess/> : <ExpandMore/>}
                    </ListItemButton>
                    <Collapse in={open} timeout="auto" unmountOnExit>
                      <List component="div" disablePadding>
                        <ListItemButton sx={{pl: 4}}>
                          <ListItemIcon>
                            <StarBorder/>
                          </ListItemIcon>
                          <ListItemText primary="Starred"/>
                        </ListItemButton>
                      </List>
                    </Collapse>
                  </List></>


              )
            }
          }
        })}
      </TabPanel>
      <TabPanel value={value.toString()} index="three">
        Content for Item Three
      </TabPanel>
    </Box>
  );
}