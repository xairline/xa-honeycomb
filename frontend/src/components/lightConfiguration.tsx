import Accordion from '@mui/material/Accordion';
import Card from '@mui/material/Card';
import Typography from '@mui/material/Typography';
import * as React from 'react';
import {pkg} from "../../wailsjs/go/models";
import {
  AccordionDetails,
  AccordionSummary,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow
} from "@mui/material";
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import DatarefValue from './datarefValue';

interface LightsProps {
  title: string;
  lights?: pkg.Leds;
  keys: string[];
}

export default function LightConfiguration(props: LightsProps) {
  const [open, setOpen] = React.useState(true);

  const handleClick = () => {
    setOpen(!open);
  };

  return (
    <Card variant="outlined">

      {/*<CardContent>*/}

      <Accordion>
        <AccordionSummary
          expandIcon={<ExpandMoreIcon/>}
          aria-controls="panel1-content"
          id="panel1-header"
        >
          <Typography variant="h6" component="div" sx={{textAlign: 'left'}}>
            {props.title}
          </Typography>
        </AccordionSummary>
        <AccordionDetails>

        </AccordionDetails>

        {Object.keys(props.lights || {}).map((key) => {
          if (!props.keys.includes(key)) {
            return null;
          }

          return (
            <>
              <AccordionSummary
                expandIcon={<ExpandMoreIcon/>}
                aria-controls="panel1-content"
                id="panel1-header"
              >
                <Typography variant="h6" component="div" sx={{textAlign: 'left'}}>
                  {key.toUpperCase()}
                </Typography>
              </AccordionSummary>
              <AccordionDetails>
                {
                  <TableContainer component={Card}>
                    <Table sx={{marginTop: "24px"}} size="small" aria-label="a dense table">
                      <TableHead>
                        {/*<TableRow>*/}
                        {/*  /!*<TableCell>LED</TableCell>*!/*/}
                        {/*  <TableCell align="right">Dataref</TableCell>*/}
                        {/*  <TableCell align="right">Operator</TableCell>*/}
                        {/*  <TableCell align="right">Threshold</TableCell>*/}
                        {/*  <TableCell align="right">Index</TableCell>*/}
                        {/*</TableRow>*/}
                      </TableHead>
                      <TableBody>
                        {
                          // @ts-ignore
                          props.lights?.[key]?.datarefs?.map((dataref) => {
                            return (
                              <TableRow
                                key={key.toUpperCase()}
                                sx={{'&:last-child td, &:last-child th': {border: 0},}}
                              >
                                {/*<TableCell component="th" scope="row">*/}
                                {/*  {key.toUpperCase()}*/}
                                {/*</TableCell>*/}
                                <TableCell align="left" sx={{width: "300px"}}>{dataref.dataref_str}</TableCell>
                                <TableCell align="left" sx={{width: "40px"}}>{dataref.operator}</TableCell>
                                <TableCell align="left" sx={{width: "40px"}}>{dataref.threshold || "0"}</TableCell>
                                <TableCell align="left" sx={{width: "40px"}}>{dataref.index || "0"}</TableCell>
                                <TableCell align="left" sx={{width: "40px"}}><DatarefValue dataref={dataref.dataref_str}
                                                                                           index={dataref.index || 0}/></TableCell>
                              </TableRow>
                            )
                          })
                        }
                      </TableBody>
                    </Table>
                  </TableContainer>
                }
              </AccordionDetails>
            </>
          )
        })
        }

      </Accordion>
      {/*</CardContent>*/
      }
    </Card>
  )
    ;
}