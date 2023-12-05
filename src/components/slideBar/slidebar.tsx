import React, { useState } from "react";
import "./slidestyle.css";
import { Link } from "react-router-dom";
import { CaretDownOutlined } from '@ant-design/icons';
import Logo from '../logo/logoBook';
import Logofront from '../logoF/logofront'

type MenuState = Record<string, boolean>;

export const Slidebar = (): JSX.Element => {
    const [menuVisible, setMenuVisible] = useState(false);

    const toggleMenu = (buttonName: string) => {
        // Implement logic to toggle menu visibility based on button clicked
        if (buttonName === 'button5') {
            setMenuVisible(!menuVisible);
        }
        // Implement logic for other buttons if needed
    };
   
    return (
            <div  className="log">
                <div className="logo1">
                    <Logo></Logo>
                </div>
                <div className="textlogo" id="position011">
                    <Logofront></Logofront>
                </div>
                    {/* Button 5 */}
                    <button /*{onClick={}}*/ className="menubutton1"  >
                        การยืมหนังสือ
                    </button>
                    {/* Button 4 */}
                    <button  /*{onClick={}}*/ className="menubutton1"  >
                        การคืนหนังสือ 
                    </button>
                    {/* Button 3 */}
                    <button  /*{onClick={}}*/ className="menubutton1"  >
                        หนังสือที่ได้รับบริจาค
                    </button>
                    {/* Button 2 */}
                    <button  /*{onClick={}}*/ className="menubutton1" >
                        คำขอหนังสือ
                    </button>
                    {/* Replicate for buttons 3, 4, and 5 */}
                    <button  className="menubutton1" onClick={() => toggleMenu('button5')} >
                        จัดการ <CaretDownOutlined />
                    </button>
                    {menuVisible && (
                        <div className="navigation">
                            <Link to="/page1">รายการหนังสือ</Link>
                            <Link to="/page2">รายการของรางวัล</Link>
                            <Link to="/page3">ห้องอ่านหนังสือ</Link>
                        </div>
                    )}

                    
            </div>
        
    
    );
  };
  export default Slidebar;