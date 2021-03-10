import React from 'react';
import PageTypes from '../../Constants/PageTypes';
import MainPageContent from './Content/MainPageContent/MainPageContent';
import SignOutButton from './Components/SignOutButton/SignOutButton';
import UpdateName from './Components/UpdateName/UpdateName';

const Main = ({ page, setPage, setAuthToken, plants, setUser, user, setPlants, addPlantCallback, toggleModal }) => {
    let content;
    let contentPage = true;
    switch (page) {
        case PageTypes.signedInMain:
            content = <MainPageContent user={user} setPage={setPage} plants={plants} setPlants={setPlants} addPlantCallback={addPlantCallback} toggleModal={toggleModal}/>;
            break;
        case PageTypes.signedInUpdateName:
            content = <UpdateName user={user} setUser={setUser} />;
            break;
        default:
            content = <>
            Error, invalid path reached
            {contentPage && <button onClick={(e) => setPage(e, PageTypes.signedInMain)}>Back to main</button>}</>;
            break;
    }
    return <>
        {content}
        <SignOutButton setUser={setUser} setAuthToken={setAuthToken} />
    </>
}

export default Main;