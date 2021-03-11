import React, { useState, useEffect } from 'react';
import PageTypes from '../../../../Constants/PageTypes';
import api from '../../../../Constants/APIEndpoints';
import {PlantList} from '../../Components/PlantList.js'
import {AddPlantModal} from '../../Components/AddPlant.js'


const MainPageContent = ({ user, plants, setPage, addPlantCallback, toggleModal }) => {
    //const [plant, newPlant] = useState(null)
    
    // async function fetchPlant() {
    //     const response = await fetch(api.base + api.handlers.myuserPlant, {
    //         method: "GET",
    //         headers: new Headers({
    //             "Authorization": localStorage.getItem("Authorization")
    //         })
    //     });
    //     if (response.status >= 300) {
    //         // const error = await response.text();
    //         //Probably need to take in more information so it can add the watering schedule?
    //         newPlant(user.photoURL)
    //         return;
    //     }

    //     //not sure what this blob thing is
    //     const imgBlob = await response.blob();
    //     newPlant(URL.createObjectURL(imgBlob));
    // }

    // useEffect(() => {
    //     fetchPlant();
    //     return;
    // }, []);

    return <>
      <header>
        <nav className="navbar">
          <span><h1 className="navbar-brand">Plant Tracker</h1></span>
          <AddPlantModal addPlantCallback={addPlantCallback} toggleModal={toggleModal} isModalOpen={false}></AddPlantModal>        </nav>
      </header>
      <PlantList plants={plants}/>

      <div><button onClick={(e) => { setPage(e, PageTypes.signedInAddedPlant) }}>Add Plant</button></div>

    </>
}


export default MainPageContent;