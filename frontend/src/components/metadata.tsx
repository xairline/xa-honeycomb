import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import * as React from 'react';
import {pkg} from "../../wailsjs/go/models";
import CheckCircleIcon from '@mui/icons-material/CheckCircle';

interface ListProps {
  metadata?: pkg.Metadata;
}

export default function Metadata(props: ListProps) {
  const [open, setOpen] = React.useState(true);

  const handleClick = () => {
    setOpen(!open);
  };

  return (
    <Card variant="outlined">

      <CardContent>

        <Typography variant="h5" component="div">
          {props.metadata?.name}
        </Typography>
        <Typography gutterBottom sx={{color: 'text.secondary', fontSize: 14}}>
          {props.metadata?.description}
        </Typography>
        {props.metadata?.enabled ? <CheckCircleIcon sx={{color: 'forestgreen', fontSize: 25}}/> : null}
      </CardContent>
    </Card>
  );
}