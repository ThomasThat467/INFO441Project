// import React, { useState, useEffect } from 'react';
// import PageTypes from '../../../../Constants/PageTypes';
// import api from '../../../../Constants/APIEndpoints';
import {PlantList} from '../../Components/PlantList.js'
import {AddPlantModal} from '../../Components/AddPlant.js'


const MainPageContent = ({ plants, addPlantCallback, getCurrentPlants, toggleModal, setUser, setAuthToken}) => {

    return <>
      <header>
        <nav className="navbar">
          <div>
            <h1>Plant Tracker</h1>
          </div>
          <AddPlantModal addPlantCallback={addPlantCallback} getCurrentPlants={getCurrentPlants} toggleModal={toggleModal} isModalOpen={false} setUser={setUser} setAuthToken={setAuthToken}></AddPlantModal>
        </nav>
      </header>
      <body>
          <PlantList plants={plants} getCurrentPlants={getCurrentPlants}/>        
      </body>

    </>
}


export default MainPageContent;