import React, {  useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import {
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
} from '@material-ui/core';
import { Alert, AlertTitle } from '@material-ui/lab';
import AddCircleOutlinedIcon from '@material-ui/icons/AddCircleOutlined';
import CancelRoundedIcon from '@material-ui/icons/CancelRounded';
import InputAdornment from '@material-ui/core/InputAdornment';
import PersonIcon from '@material-ui/icons/Person';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
import { EntDepartment, EntJobtitle, EntPersonal, EntGender } from '../../api';

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
      margin: theme.spacing(2),
   },
   insideLabel: {
    margin: theme.spacing(1),
  },
   button: {
    marginLeft: '40px',
  },
   textField: {
    width: 500 ,
    marginLeft:7,
    marginRight:-7,
   },
    select: {
      width: 500 ,
      marginLeft:7,
      marginRight:-7,
      //marginTop:10,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
  }),
);


export default function CreateUser() {
  const classes = useStyles();
  const http = new DefaultApi();

  const [personals, setPersonal] = React.useState<EntPersonal[]>([]);
  
  const [jobtitles, setJobtitles] = React.useState<EntJobtitle[]>([]);
  const [departments, setDepartments] = React.useState<EntDepartment[]>([]);
  const [genders, setGenders] = React.useState<EntGender[]>([]);
  
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  const [department, setDepartment] = useState(Number);
  const [jobtitle, setJobtitle] = useState(Number);
  const [gender, setGender] = useState(Number);

  const [personalName, setPersonalName] = useState(String);


  useEffect(() => {
    const getDepartments = async () => {
      const res = await http.listDepartment({ limit: 10, offset: 0 });
      setLoading(false);
      setDepartments(res);
      console.log(res);
    };
    getDepartments();

    const getJobtitles = async () => {
      const res = await http.listJobtitle({ limit: 10, offset: 0 });
      setLoading(false);
      setJobtitles(res);
      console.log(res);
    };
    getJobtitles();

    const getGenders = async () => {
      const res = await http.listGender({ limit: 10, offset: 0 });
      setLoading(false);
      setGenders(res);
      console.log(res);
    };
    getGenders();

  }, [loading]);

  const getPersonal = async () => {
    const res = await http.listPersonal({ limit: 10, offset: 0 });
    setPersonal(res);
  };

  const handlePersonalNameChange = (event: any) => {
    setPersonalName(event.target.value as string);
  };
  
  const JobtitlehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setJobtitle(event.target.value as number);
  };

  const DepartmenthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDepartment(event.target.value as number);
  };

  const GenderhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setGender(event.target.value as number);
  };


// create personal
const CreatePersonal = async () => {
  const personal = {
    personalName : personalName,
    jobtitle : jobtitle,
    department : department,
    gender : gender,
  };
  console.log(personals);
  const res: any = await http.createPersonal({ personal: personal });
  console.log("bruhhhhhhhhh");
  setStatus(true);
  
  if (res.id != '') {
    setAlert(true);
  } else {
    setAlert(false);
  }
  const timer = setTimeout(() => {
    setStatus(false);
  }, 3000);
};

return (
  <Page theme={pageTheme.tool}>

    <Header
      title={`INPATIENT SYSTEM`}
      type="Personal System"> 
    </Header>

    <Content>
      <ContentHeader title="Personal System"> 
            <Button onClick={() => {CreatePersonal();}} variant="contained"  color="primary" startIcon={<AddCircleOutlinedIcon/>}> Create new personal </Button>
            <Button style={{ marginLeft: 20 }} component={RouterLink} to="/table" variant="contained" startIcon={<CancelRoundedIcon/>}>  Dismiss </Button>
      </ContentHeader>  
      <div className={classes.root}>
        <form noValidate autoComplete="off">
          <FormControl
            //fullWidth
            //className={classes.margin}
            variant="outlined"
          >
             <div className={classes.paper}><strong>ชื่อ-นามสกุล</strong></div>
            <TextField className={classes.textField}
  //          style={{ width: 500 ,marginLeft:7,marginRight:-7}}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <PersonIcon />
                </InputAdornment>
              ),
            }}
              id="personalName"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={personalName}
              onChange={handlePersonalNameChange}
            />

            <div className={classes.paper}><strong>สายงาน</strong></div>
            <Select className={classes.select}
              //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="secondary"
              labelId="jobtitle-label"
              id="jobtitle"
              value={jobtitle}
              onChange={JobtitlehandleChange}
            >
              <InputLabel className={classes.insideLabel} id="faculty-label">เลือกสายงาน(Jobtitle)</InputLabel>

              {jobtitles.map((item: EntJobtitle) => (
                <MenuItem value={item.id}>{item.jobtitlename}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>แผนก</strong></div>
            <Select className={classes.select}
              //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="secondary"
              id="department"
              value={department}
              onChange={DepartmenthandleChange}
            >
              <InputLabel className={classes.insideLabel}>เลือกแผนก(Department)</InputLabel>

              {departments.map((item: EntDepartment) => (
                <MenuItem value={item.id}>{item.departmentname}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>เพศ</strong></div>
            <Select className={classes.select}
              //style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="secondary"
              id="gender"
              value={gender}
              onChange={GenderhandleChange}
            >
              <InputLabel className={classes.insideLabel}>เลือกเพศ(Systemmember)</InputLabel>

              {genders.map((item: EntGender) => (
                <MenuItem value={item.id}>{item.gendername}</MenuItem>
              ))}
            </Select>

            {status ? ( 
                    <div className={classes.margin} style={{ width: 500 ,marginLeft:30,marginRight:-7,marginTop:16}}>
            {alert ? ( 
                    <Alert severity="success"> <AlertTitle>Success</AlertTitle> Complete data — check it out! </Alert>) 
            : (     <Alert severity="warning"> <AlertTitle>Warining</AlertTitle> Incomplete data — please try again!</Alert>)} </div>
          ) : null}
          
          </FormControl>

        </form>
      </div>
    </Content>
  </Page>
);
}
