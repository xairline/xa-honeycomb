import * as React from 'react';
import {useEffect} from 'react';
import Box from '@mui/material/Box';
import {pkg} from '../../wailsjs/go/models';
import {Divider, Grid2, IconButton, List, ListItemButton, ListItemIcon, Switch, Typography} from "@mui/material";
import ListItemText from '@mui/material/ListItemText';
import AddIcon from '@mui/icons-material/Add';

interface TabPanelProps {
  profiles: pkg.Profile[];
  selectedIndex: number;
  handleListItemClick: (event: React.MouseEvent<HTMLDivElement, MouseEvent>, index: number) => void;
}


export default function Profiles(props: TabPanelProps) {
  const [checked, setChecked] = React.useState([] as string[]);

  const handleToggle = (value: string) => () => {
    const currentIndex = checked.indexOf(value);
    const newChecked = [...checked];

    if (currentIndex === -1) {
      newChecked.push(value);
    } else {
      newChecked.splice(currentIndex, 1);
    }

    setChecked(newChecked);
  };

  useEffect(() => {
    const newChecked = props.profiles
      .filter((profile) => profile.metadata?.enabled)
      .map((profile) => profile.metadata?.name || ""); // Collect names of enabled profiles
    setChecked(newChecked);
  }, [props.profiles]);

  return (
    <Box sx={{width: '100%', overflow: "hidden", bgcolor: "#022e35", height: "100vh"}}>
      < List
        component="nav"
        subheader={
          <>
            <Box>
              <Grid2 container spacing={2}>
                <Grid2 size={4}>
                </Grid2>
                <Grid2 size={4}>
                  <Typography variant="h5" style={{marginTop: "12px", marginBottom: "12px"}}>Profiles</Typography>
                </Grid2>
                <Grid2 size={4}>
                  <IconButton
                    onClick={() => {
                      alert("TODO: Add new profile")
                    }}
                    style={{marginTop: "12px", backgroundColor: "#8EAF91FF"}}
                    size="small"

                  ><AddIcon fontSize="small"></AddIcon></IconButton>
                </Grid2>


              </Grid2>

            </Box>
            <Divider/>
          </>
        }
      >
        {
          props.profiles.map((profile, index) => {
            return (
              <ListItemButton
                selected={props.selectedIndex === index}
                onClick={(event) => props.handleListItemClick(event, index)}
              >
                <ListItemIcon>
                  <Switch
                    edge="end"
                    onChange={handleToggle(profile.metadata?.name || "")}
                    checked={
                      checked.includes(profile.metadata?.name || "asdf")
                    }
                    inputProps={{
                      'aria-labelledby': 'switch-list-label-wifi',
                    }}
                    disabled={true}
                  />
                </ListItemIcon>
                <ListItemText primary={profile.metadata?.name} style={{marginLeft: "12px"}}/>
              </ListItemButton>
            )
          })
        }
      </List>

    </Box>
  )
    ;
}