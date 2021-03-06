import React from 'react';
import PageTypes from '../../Constants/PageTypes';
import MainPageContent from './Content/MainPageContent/MainPageContent';

const Main = ({ page, setPage, setAuthToken, plants, setUser, user, setPlants, addPlantCallback, getCurrentPlants, toggleModal }) => {
    let content;
    let contentPage = true;
    if (page === PageTypes.signedInMain){
      content = <MainPageContent user={user} setPage={setPage} plants={plants} setPlants={setPlants} addPlantCallback={addPlantCallback} getCurrentPlants={getCurrentPlants} toggleModal={toggleModal} setUser={setUser} setAuthToken={setAuthToken}/>;
    } else {
      content = <>
      Error, invalid path reached
      {contentPage && <button onClick={(e) => setPage(e, PageTypes.signedInMain)}>Back to main</button>}</>;
    }
    return <>
        {content}
    </>
}

export default Main;