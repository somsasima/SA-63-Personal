import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import DeleteIcon from '@material-ui/icons/Delete';
import {
  Content,
  Header,
  Page,
  pageTheme,
} from '@backstage/core';
import PersonAddRoundedIcon from '@material-ui/icons/PersonAddRounded';
import HomeRoundedIcon from '@material-ui/icons/HomeRounded';
import { EntPersonal } from '../../api';

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
  table: {
    minWidth: 650,
  },
  buttonRight: {
    marginLeft: theme.spacing(150),
    marginBottom: theme.spacing(2),
  },
  }),
);
 
export default function ComponentsTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [personals, setPersonals] = useState<EntPersonal[]>([]);
  const [loading, setLoading] = useState(true);
  

  // get personals
  useEffect(() => {
    const getPersonals = async () => {
      const res:any = await http.listPersonal({ limit: 10, offset: 0 });
      setLoading(true);
      setPersonals(res);
      console.log(res);
    };
    getPersonals();
  }, [loading]);
  
  // delete button
  const deletePersonals = async (id: number) => {
    const res = await http.deletePersonal({ id: id });
    setLoading(true);
  };

    // clear input form
    function clear() {
      setPersonals([]);
    }
  
 
  // ui 
 return (
  <Page theme={pageTheme.tool}>
    <Header title={`Personal System`} type="INPATIENT SYSTEM" >
    <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="default" href="/user" startIcon={<PersonAddRoundedIcon />}> New personal</Button>
    <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="primary" href="/" startIcon={<HomeRoundedIcon/>}> home</Button>
    </Header>
    
    <Content>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No</TableCell>
           <TableCell align="center">PersonalName</TableCell>
           <TableCell align="center">Jobtitle</TableCell>
           <TableCell align="center">Department</TableCell>
           <TableCell align="center">Gender</TableCell>
           <TableCell align="center">Mange</TableCell>
         </TableRow>
       </TableHead>

       <TableBody>
         {personals.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.personalName}</TableCell>
             <TableCell align="center">{item.edges?.jobtitle?.jobtitlename}</TableCell>
             <TableCell align="center">{item.edges?.department?.departmentname}</TableCell>
             <TableCell align="center">{item.edges?.gender?.gendername}</TableCell>
             <TableCell align="center">
             <Button
                 onClick={() => {
                   deletePersonals(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
                 startIcon={<DeleteIcon/>}
                 href="/table"
               >
                 Delete
               </Button>
 
             </TableCell>

           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Content>
  </Page>
);
}
