import React, { FC, useEffect,useState } from 'react';

import Swal from 'sweetalert2'; // alert

import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,

} from '@backstage/core';


import PersonIcon from '@material-ui/icons/Person';
import { EntSubject } from '../../api/models/EntSubject'; //import ข้อมูลจาก backend
import { EntDegree } from '../../api/models/EntDegree'; 
import { EntDepartment } from '../../api/models/EntDepartment'; 
import { DefaultApi} from '../../api/apis';

import {

  
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
} from '@material-ui/core';


const useStyles = makeStyles((theme: Theme) =>

  createStyles({
    root: {
      '& .MuiTextField-root': {
        marginLeft: theme.spacing(9),
        width: '30ch',
      },
    },

    formControl: {
      width:300,

    },

    text: {  //ชื่อชมรม
        marginLeft: theme.spacing(5),
        marginRight: theme.spacing(4),
        marginTop: theme.spacing(1),
        fontSize: 20,
    },
    text2: {
      marginLeft: theme.spacing(4),
      marginTop: theme.spacing(4),
      fontSize: 20,
  },
  text3: {
    marginLeft: theme.spacing(4),
    marginTop: theme.spacing(4),
    fontSize: 20,
},

  box1: {
    marginLeft: theme.spacing(7),
    marginTop: theme.spacing(2),
    width: '30ch',  
},

  box2: {
    marginLeft: theme.spacing(5),
    marginTop: theme.spacing(4),
    width: '30ch',  
  },

  box3: {
    marginLeft: theme.spacing(3),
    marginTop: theme.spacing(2),
    width: '30ch',  
},


  textfill1: {  //ชื่อกิจกรรม
    marginLeft: theme.spacing(10),
    marginTop: theme.spacing(3),
    width: '80ch',
  },
  textfill2: {
    marginLeft: theme.spacing(10),
    marginTop: theme.spacing(4),
    width: '80ch',
  },
  textfill3: {
    marginLeft: theme.spacing(14),
    marginTop: theme.spacing(4),
    width: '80ch',
  },
  textfill4: {
    marginLeft: theme.spacing(16),
    marginTop: theme.spacing(4),
    width: '30ch',
  },
  button: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexWrap: 'wrap',
    justifyContent: 'center',
  },
  text4: {
    marginLeft: theme.spacing(2),
    marginTop: theme.spacing(1),
    marginBottom:  theme.spacing(1),
    fontSize: 14,
},
text5: {
  marginLeft: theme.spacing(2),
  marginTop: theme.spacing(1),
  marginBottom:  theme.spacing(1),
  fontSize: 14,
  color: "blue",
},
  }),
);

//โครงสร้างของ interface
interface currency{
    DepartmentID :      number;
	SubjectID :      number;
	DegreeID  :      number;
	CourseName : string;
	
}

const WelcomePage: FC<{}> = () => {
const classes = useStyles();

const api = new DefaultApi();

 const [currency, setCurrency] = React.useState<Partial<currency>>({});
 const [departments, setDepartments ] = React.useState<EntDepartment[]>([]);
 const [subjects, setSubjects ] = React.useState<EntSubject[]>([]);
 const [degrees, setDegrees ] = React.useState<EntDegree[]>([]);
 const Toast = Swal.mixin({  //ตั้งค่าการแจ้งเตือน
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: toast => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});

const handleChange = ( event: React.ChangeEvent<{ name?: string; value: unknown }>,) =>{
    const name = event.target.name as keyof typeof WelcomePage;
    const { value } = event.target;
    setCurrency({ ...currency, [name] : value });
    console.log(currency);
};


 const getDepartments = async () => {
  const res = await api.listDepartment({ limit: 10, offset:0});
  setDepartments(res);
 };

 const getSubjects = async () => {//ดึงข้อมูล
  const res = await api.listSubject({ limit: 10, offset:0}); //ดึงจาก listUser
  setSubjects(res); //ส่งค่าไป user ข้างบน
 };

 const getDegrees = async () => {
  const res = await api.listDegree({ limit: 10, offset:0});
  setDegrees(res);
 };

 
 useEffect(() => { //เรียกใช้
    getDepartments()
    getSubjects()
    getDegrees()
 }, []
 );

 // clear input form
 function clear() {
  setCurrency({});
}

//6 กดบันทึกที่หน้า ui 
 function save() {
    const apiUrl = 'http://localhost:8080/api/v1/activitys'; //url จาก main.go + activitys ตรง crud ใน ctl
    const requestOptions = { //ส่งคำขอที่เป็นแบบ ใส่ค่าข้อมูล
      method: 'POST', // 7 เพิ่มเข้ารายการกิจกรรม (ใส่ค่าข้อมูล)
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(currency),
  };
    console.log(currency); // ดูข้อมูล กด F12 เลือก Tab Console
    fetch(apiUrl, requestOptions) //ส่งคำขอผ่าน apiUrl ให้ ctl ส่งค่ากลับมา
      .then(response => response.json())
      .then(data => {//ส่งค่า data = ac
        console.log(data);
        if (data.status === true) { //เชื่อมกับ controllerหลัก ตรง create
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }
    
  /*const Createpreempt = async () => { 
    const preemption = {
      user: 1,
      roominfo: roomid,
      purpose: purposeid,
    
      added: preempttime + ":00+07:00"
    
    };
    
    console.log(preemption);
    
    const res: any = await api.createActivity({ preemption: preemption });
    
    
    
     };*/

 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`สร้างหลักสูตร`}
       
     ></Header>

     <ContentHeader title="">
     <div><PersonIcon style={{ fontSize: 100}} color="secondary" ></PersonIcon> </div>
     <div><FormControl className={classes.text4} > <text>chanwit</text> </FormControl></div>
     <FormControl className={classes.text5} > <text>Logout</text> </FormControl>
    </ContentHeader>

     <Content>

     <div>
       <FormControl className={classes.text} > <text>ชื่อสำนักวิชา</text> </FormControl>   
       <FormControl fullWidth  variant="outlined" 
        className={classes.box3}>
        <InputLabel>select</InputLabel>
        <Select
          name="DepartmentID" //เก็บไว้ตัวแปรไหน
          value={currency.DepartmentID || ''}  //เก็บไว้ที่ไหน
          onChange={handleChange}
        >
        {departments.map(item => { //ฟังชันไว้แสดงข้อมูลใน  user combo box
            return (
              <MenuItem key ={item.id} value = {item.id}>
                {item.uSERNAME}
                </MenuItem>
            );
          }  
            )}
          
        </Select>
        </FormControl>
      
      </div> 

      
       

     <div>
       <FormControl className={classes.text} > <text>ชื่อชมรม</text> </FormControl>   
       <FormControl fullWidth  variant="outlined" 
        className={classes.box1}>
        <InputLabel>select</InputLabel>
        <Select
          name="CLUBID"
          value={currency.CLUBID || ''}
          onChange={handleChange}
        >
        {clubs.map(item => {
            return (
              <MenuItem key ={item.id} value = {item.id}>
                {item.cLUBNAME}
                </MenuItem>
            );
          }  
            )}
          
        </Select>
        </FormControl>
      
      </div> 



       <div>
       <FormControl className={classes.text3} ><text >ประเภทกิจกรรม</text> </FormControl>
       <FormControl fullWidth  variant="outlined" 
        className={classes.box2}>
        <InputLabel>select</InputLabel>
        <Select
          name="TYPEID"
          value={currency.TYPEID || ''}
          onChange={handleChange}
        >
        {kinds.map(item => {
            return (
              <MenuItem key ={item.id} value = {item.id}>
                {item.tYPENAME}
                </MenuItem>
            );
          }  
            )}
          
        </Select>
        </FormControl>
      
      </div>    

      <div>
      <FormControl className={classes.text3} >  <text >ชื่อกิจกรรม</text> </FormControl>
       <FormControl className={classes.textfill1} >
           <TextField id="outlined-basic"  variant="outlined" 
           name="ACTIVITYNAME"
           defaultValue="พิมพ์ชื่อกิจกรรม"
           value = {currency.ACTIVITYNAME}
           onChange={handleChange}
           />
      </FormControl>
      </div>
      
      

     
     
      
      

      <div className={classes.button}>
      <Button size="large" variant="contained" color="secondary" onClick = {save}> บันทึกกิจกรรม </Button>
      </div>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;