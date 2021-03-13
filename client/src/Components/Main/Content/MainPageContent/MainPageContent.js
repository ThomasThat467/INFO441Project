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
        <div className="content-inside">
          <PlantList plants={plants} getCurrentPlants={getCurrentPlants}/>        
        </div>
      </body>
      <footer>
          <p className="footer-text">&#169; Hailey Meister, Jisu Kim, Eric Gabrielson, and Thomas That</p>
      </footer>
    </>
}


export default MainPageContent;