import * as React from 'react';
import List from '@mui/material/List';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import Collapse from '@mui/material/Collapse';
import ExpandLess from '@mui/icons-material/ExpandLess';
import ExpandMore from '@mui/icons-material/ExpandMore';
import StarBorder from '@mui/icons-material/StarBorder';
import {pkg} from "../../wailsjs/go/models";
import {Badge, Divider} from "@mui/material";
import {LightbulbOutlined} from "@mui/icons-material";
import Dataref = pkg.Dataref;
import Command = pkg.Command;

interface ListProps {
  profile: pkg.Profile;
}

export default function NestedList(props: ListProps) {
  const [open, setOpen] = React.useState(true);

  const handleClick = () => {
    setOpen(!open);
  };

  return (
    <List
      sx={{width: '100%'}}
      component="nav"
      aria-labelledby="nested-list-subheader"
    >
      {Object.keys(props.profile).map((key: string) => {
        if (key === "metadata") {
          return null;
        }

        return (
          <>
            <ListItemButton onClick={handleClick}>
              <ListItemIcon>
                <Badge
                  color="success"
                  variant="dot"
                  sx={{
                    '& .MuiBadge-dot': {
                      width: 10,  // Customize dot width
                      height: 10, // Customize dot height
                      borderRadius: '50%', // Ensure it's circular
                    },
                  }}
                >
                  <LightbulbOutlined/>
                </Badge>
              </ListItemIcon>
              <ListItemText primary={key.toUpperCase()}/>
              {open ? <ExpandLess/> : <ExpandMore/>}
            </ListItemButton>
            <Collapse in={open} timeout="auto" unmountOnExit>
              <List component="div" disablePadding>
                {

                  // @ts-ignore
                  props.profile[key]?.datarefs?.map(
                    (dataref: Dataref) => {
                      return (
                        <ListItemButton sx={{pl: 8}}>
                          <ListItemIcon>
                            <StarBorder/>
                          </ListItemIcon>
                          <ListItemText primary={
                            // @ts-ignore
                            `${dataref.dataref_str} ${dataref.operator} ${dataref.threshold}`
                          }/>

                        </ListItemButton>
                      )
                    }
                  )

                }
                {

                  // @ts-ignore
                  props.profile[key]?.commands?.map(
                    (command: Command) => {
                      return (
                        <ListItemButton sx={{pl: 8}}>
                          <ListItemIcon>
                            <StarBorder/>
                          </ListItemIcon>
                          <ListItemText primary={
                            // @ts-ignore
                            command.command_str
                          }/>
                        </ListItemButton>
                      )
                    }
                  )

                }
              </List>
            </Collapse>
            <Divider/>

          </>
        )
      })}
    </List>
  );
}