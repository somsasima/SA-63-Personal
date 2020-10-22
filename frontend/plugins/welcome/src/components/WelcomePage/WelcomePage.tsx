import React, { FC } from 'react';
import Button from '@material-ui/core/Button';
import AddBoxRoundedIcon from '@material-ui/icons/AddBoxRounded';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import {
 Content,
 Header,
 Page,
 pageTheme,
} from '@backstage/core';
import {
  Avatar,
} from '@material-ui/core';

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(70),
    },
  }),
);

const WelcomePage: FC<{}> = () => {
  const classes = useStyles();
 return (
   
   <Page theme={pageTheme.tool}>
     <Header
       title="INPATIENT SYSTEM" type="G05">
     </Header>

      <div className={classes.paper}>
        <Content>
            <Button 
              variant="outlined" 
              color="secondary" 
              href="/table" 
              startIcon={<AddBoxRoundedIcon />}> 
              PERSONAL SYSTEM
            </Button>
        </Content>
      </div>
   </Page>
 );
};

export default WelcomePage;