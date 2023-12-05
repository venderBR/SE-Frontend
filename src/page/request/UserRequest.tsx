import React from "react";
import "./style.css";
import { Link } from "react-router-dom";
import { Image } from 'antd';
import { Upload,Progress,Form,message,Button,Input } from 'antd';
import {Navbar} from '../../components/navBar/navbar';

export const Request = (): JSX.Element => {
   
    return (
        <div className="element1">
            <div className="background"> 
                <div className="blockbar"> 
                    <Navbar></Navbar>
                </div>
                <div className="block">
                    <div className="text" id="position">ชื่อหนังสือ :  </div> 
                    <div className="text" id="position1">ชื่อผู้เขียน :  </div>
                    <div className="text" id="position2">ชื่อสำนักพิมพ์ :  </div>
                    <div className="text" id="position3">แสดงเหตุผล :  </div>
                    <Form style={{height:"100%"}}>
                        <Form.Item  name="Title">
                            <Input  placeholder="ชื่อหนังสือ" />
                        </Form.Item>
                        <Form.Item name="Writer">
                            <Input placeholder="ชื่อผู้เขียน" />
                        </Form.Item>
                        <Form.Item  name="Publisher">
                            <Input  placeholder="ชื่อสำนักพิมพ์" />
                        </Form.Item>
                        <Form.Item style={{marginBottom:"4.5%"}} name="Reson">
                            <Input  placeholder="แสดงเหตุผล" />
                        </Form.Item>
                        <Form.Item  name="block1" >
                            <Button type="primary" htmlType="submit"  className="button" >SUBMIT</Button>
                        </Form.Item>
                        
                    </Form>
                
                </div>
                
                <div className="blocktext"> 
                    <img className="search" 
                        alt="Search" 
                        src="https://c.animaapp.com/OFw4BeUY/img/search-1@2x.png" 
                    />
                    <div className="text" id="frist">Request for New Books   </div>
                    <div className="text" id="song">in the Library  </div>
                    <div className="text" id="three">
                        " เราต้องการคำแนะนำเกี่ยวกับหนังสือที่คุณชื่นชอบอ่าน!
                        <br />มาช่วยแบ่งปันความชื่นชอบของคุณและเพื่อแนวทาง
                        <br />สำหรับการสร้างคอลเล็กชันหนังสือที่หลากหลาย
                        <br />เพื่อเสริมสร้างห้องสมุดของเรา
                        <br />ให้เต็มไปด้วยเล่มที่คุณต้องการอ่าน 
                        <br /><br />
                        คำแนะนำของคุณจะช่วยเราได้มาก<br />เพื่อเลือกหนังสือที่สอดคล้องกับความสนใจของคุณ! "
                    </div>
                    <img className="pngtreeround" 
                        alt="Pngtreeround" 
                        src="https://c.animaapp.com/OFw4BeUY/img/-pngtree-round-creative-library-bookshelf-5304269-1.png" 
                    />
                </div>
                <div className="local" ><Link to="/donate" className="text" id="edit2" >Donate</Link></div>
                <div className="local1" ><Link to="/request" className="text" id="edit2" >Request</Link></div>
            </div>
        </div>
    );
  };
  export default Request;