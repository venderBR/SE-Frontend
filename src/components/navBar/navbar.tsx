import React, { useState } from "react";
import "./navstyle.css";
import { Link } from "react-router-dom";

export const Navbar = (): JSX.Element => {
    return (
        <div className="blockbar1"> 
            <div id="home">
                <Link to="/home"  className="text1">HOME</Link>  
            </div>
            <div id="book"  >
                <Link to="/Book" className="text1">BOOK</Link>   
            </div>
            <div id="room"  >
                <Link to="/Room" className="text1">ROOM</Link> 
            </div>
            <div id="rewards"  >
                <Link to="/Rewaeds" className="text1">REWARDS</Link>  
            </div>
            <div className="ellipse"></div>
            <div className="line"></div>
        </div>
    );
  };
  export default Navbar;