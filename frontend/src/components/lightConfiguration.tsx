import Accordion from '@mui/material/Accordion';
import Card from '@mui/material/Card';
import Typography from '@mui/material/Typography';
import * as React from 'react';
import {pkg} from "../../wailsjs/go/models";
import {
  AccordionDetails,
  AccordionSummary,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow
} from "@mui/material";
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';

interface LightsProps {
  title: string;
  lights?: pkg.Leds;
}

export default function LightConfiguration(props: LightsProps) {
  const [open, setOpen] = React.useState(true);

  const handleClick = () => {
    setOpen(!open);
  };

  function createData(
    name: string,
    calories: number,
    fat: number,
    carbs: number,
    protein: number,
  ) {
    return {name, calories, fat, carbs, protein};
  }

  const rows = [
    createData('Frozen yoghurt', 159, 6.0, 24, 4.0),
    createData('Ice cream sandwich', 237, 9.0, 37, 4.3),
    createData('Eclair', 262, 16.0, 24, 6.0),
    createData('Cupcake', 305, 3.7, 67, 4.3),
    createData('Gingerbread', 356, 16.0, 49, 3.9),
  ];

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


          return (
            <>
              {
                <TableContainer component={Paper}>
                  <Table sx={{marginTop: "24px"}} size="small" aria-label="a dense table">
                    <TableHead>
                      <TableRow>
                        <TableCell>LED</TableCell>
                        <TableCell align="right">Dataref</TableCell>
                        <TableCell align="right">Operator</TableCell>
                        <TableCell align="right">Threshold</TableCell>
                        <TableCell align="right">Index</TableCell>
                      </TableRow>
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
                              <TableCell component="th" scope="row">
                                {key.toUpperCase()}
                              </TableCell>
                              <TableCell align="right">{dataref.dataref_str}</TableCell>
                              <TableCell align="right">{dataref.operator}</TableCell>
                              <TableCell align="right">{dataref.threshold}</TableCell>
                              <TableCell align="right">{dataref.index}</TableCell>
                            </TableRow>
                          )
                        })
                      }
                    </TableBody>
                  </Table>
                </TableContainer>
              }
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