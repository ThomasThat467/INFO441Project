import React, { useState, useEffect } from 'react';
import PageTypes from '../../../../Constants/PageTypes';
import api from '../../../../Constants/APIEndpoints';
import {PlantList} from '../../Components/PlantList.js'
import {AddPlantModal} from '../../Components/AddPlant.js'

const MainPageContent = ({ user, plants, setPage }) => {
    const [avatar, setAvatar] = useState(null)

    async function fetchAvatar() {
        const response = await fetch(api.base + api.handlers.myuserAvatar, {
            method: "GET",
            headers: new Headers({
                "Authorization": localStorage.getItem("Authorization")
            })
        });
        if (response.status >= 300) {
            // const error = await response.text();
            setAvatar(user.photoURL)
            return;
        }
        const imgBlob = await response.blob();
        setAvatar(URL.createObjectURL(imgBlob));
    }

    useEffect(() => {
        fetchAvatar();
        return;
    }, []);

    return <>
      <header>
        <nav className="navbar">
          <span><h1 className="navbar-brand">Plant Tracker</h1></span>
        </nav>
      </header>
      
      {/* <PlantList plants={this.state.plants}/> */}
      {/* <AddPlantModal addPlantCallback={this.addPlantCallback} toggleModal={this.toggleModal} isModalOpen={false}></AddPlantModal> */}

      {avatar && <img className={"avatar"} src={avatar} alt={`${user.firstName}'s avatar`} />}
      <div><button onClick={(e) => { setPage(e, PageTypes.signedInAddedPlant) }}>Add Plant</button></div>

    </>
}

// class Header extends Component {
//     constructor(props){
//         super(props);
//         this.addPlantCallback = this.props.addPlantCallback;
//         this.handleSignOutCallback = this.props.handleSignOutCallback;
//     }

//     render() {
//         return (
//             <header>
//                 <nav className="navbar">
//                     <span><h1 className="navbar-brand">Plant Tracker</h1></span>

//                     <AddPlantModal addPlantCallback={this.addPlantCallback} toggleModal={this.toggleModal} isModalOpen={false}></AddPlantModal>
//                     <SignOut handleSignOutCallback={this.handleSignOutCallback}></SignOut>
                    
//                 </nav>
//             </header>
//         );
//     }
// }


export default MainPageContent;