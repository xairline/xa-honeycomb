import {Typography} from '@mui/material';
import * as React from 'react';
import {useEffect} from 'react';
import {GetXplaneDataref} from "../../wailsjs/go/main/App";

interface ListProps {
  dataref: string;
  index: number;
}

export default function DatarefValue(props: ListProps) {
  const [value, setValue] = React.useState(0);

  useEffect(() => {
    // run the function to get the xplane data evey 5 seconds
    const interval = setInterval(() => {
      GetXplaneDataref(props.dataref).then((res) => {
        const obj = JSON.parse(res);
        console.log(props, obj.data);
        if (obj?.data?.length > 0) {
          setValue(obj.data[props.index]);
        } else {
          setValue(obj.data as any);
        }

      });
    }, 1000);
  }, []);

  return (
    <Typography variant="h6" sx={{fontSize: 16, color: 'aqua'}}>
      {value}
    </Typography>

  );
}