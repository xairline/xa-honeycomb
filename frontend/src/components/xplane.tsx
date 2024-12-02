import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import * as React from 'react';
import {useEffect} from 'react';
import {pkg} from "../../wailsjs/go/models";
import {GetXplane} from "../../wailsjs/go/main/App";
import {Divider, Skeleton} from "@mui/material";

interface ListProps {
  metadata?: pkg.Metadata;
}

export default function Xplane(props: ListProps) {
  const [xplaneStat, setXplaneStat] = React.useState([]);

  useEffect(() => {
    // run the function to get the xplane data evey 5 seconds
    const interval = setInterval(() => {
      GetXplane().then((res) => {
        setXplaneStat(res as any);
      });
    }, 1000);
  }, []);

  return xplaneStat.length > 0 ? (
    <>
      <Card variant="outlined" sx={{margin: "8px", height: "92%"}}>

        <CardContent>

          <Typography variant="h5">
            Current Plane
          </Typography>
          <Divider sx={{margin: "8px"}}/>
          <Typography variant="h6" sx={{fontSize: 16, color: 'text.secondary'}}>
            {atob(JSON.parse(xplaneStat[0]).data)}
          </Typography>
          <Typography variant="h6" sx={{fontSize: 16, color: 'text.secondary'}}>
            {atob(JSON.parse(xplaneStat[1]).data)}
          </Typography>

        </CardContent>
      </Card></>

  ) : (<Skeleton/>);
}