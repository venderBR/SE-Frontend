import React, { useState } from "react";
import "./style.css";
import { Link } from "react-router-dom";
import { Image } from 'antd';
import { Upload,Progress,Form,message,Button,Input } from 'antd';
import { Slidebar } from "../../components/slideBar/slidebar";


const siderStyle: React.CSSProperties = {
    flex: '0 0 18%',
    width: '18%',
    minWidth: '18%',
    maxWidth: '18%',
};
const MainStyle: React.CSSProperties = {
    flex: '0 22% 100%',
    width: '80%%',
    minWidth: '80%',
    maxWidth: '80%',
};

type MenuState = Record<string, boolean>;

export const Requested = (): JSX.Element => {
    const [menuVisible, setMenuVisible] = useState(false);

    const toggleMenu = (buttonName: string) => {
        // Implement logic to toggle menu visibility based on button clicked
        if (buttonName === 'button5') {
            setMenuVisible(!menuVisible);
        }
        // Implement logic for other buttons if needed
    };
   
    return (
        <div className="element2">
            <div style={siderStyle} className="l">
                <Slidebar></Slidebar>
            </div>
            <div style={MainStyle} className="maincontent">
                <div className="Cblock">
                </div>
            </div>

            
        </div>
        
    
    );
  };
  export default Requested;