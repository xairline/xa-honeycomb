import Accordion from '@mui/material/Accordion';
import Card from '@mui/material/Card';
import Typography from '@mui/material/Typography';
import * as React from 'react';
import {pkg} from "../../wailsjs/go/models";
import {AccordionDetails, AccordionSummary} from "@mui/material";
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';

interface LightsProps {
  title: string;
  lights?: pkg.Leds;
}

export default function Lights(props: LightsProps) {
  const [open, setOpen] = React.useState(true);

  const handleClick = () => {
    setOpen(!open);
  };
  // TODO: Add more lights
  return (
    <Card variant="outlined">

      {/*<CardContent>*/}

      <Accordion>
        <AccordionSummary
          expandIcon={<ExpandMoreIcon/>}
          aria-controls="panel1-content"
          id="panel1-header"
        >
          <Typography variant="h5" component="div" sx={{textAlign: 'left'}}>
            {props.title}
          </Typography>
        </AccordionSummary>
        <AccordionDetails>
          <Typography gutterBottom sx={{color: 'text.secondary', textAlign: 'left'}}>
            {props.lights?.alt ? "Altitude" : null}
            {
              props.lights?.alt?.datarefs &&
              props.lights?.alt?.datarefs.map((dataref) => {
                return (
                  <Typography gutterBottom sx={{color: 'text.secondary', textAlign: 'left'}}>
                    {dataref.dataref_str} {dataref.operator} {dataref.threshold}
                  </Typography>
                )
              })
            }
          </Typography>
        </AccordionDetails>
      </Accordion>
      {/*</CardContent>*/}
    </Card>
  );
}