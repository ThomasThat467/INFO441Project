import React, { Component} from 'react';
import './index.css';
import {Button} from 'reactstrap'
import 'whatwg-fetch';
import {PlantList} from './Components/PlantList.js'
import {AddPlantModal} from './Components/AddPlant.js'
//import {LoginForm} from './Components/Login.js'
import {Route, NavLink} from 'react-router-dom';


class App extends Component {
    constructor(props){
        super(props);
        this.state = {
            plants: [],
            signedIn: false, 
            isModalOpen: false
        };
        // this.toggleModal = this.toggleModal.bind(this);
    }

    componentDidMount() {
        fetch('data/inventory.json')
        .then( (response) => {
            return response.json();
        })
        .then( (data) => {
            this.setState({plants: data.plantInventory})
        })

    }

    // handleSignUp = (props) => {
    //     console.log("handleSignUp scope")
    //     console.log(props)
    //       .then((userCredential) => {
    //         let user = userCredential.user;
    //         console.log(user);
    
    //         let updatePromise = user.updateProfile({displayName: this.state.username})
    //         return updatePromise;
    //       })
    //       .then(() => {
    //         this.setState((prevState) => {
    //           let updatedUser = {...prevState.user, displayName: this.state.username}
    //           return {user: updatedUser}; //updating the state
    //         });
    //       })
    //       .catch((err) => {
    //         this.setState({errorMessage: err.message});
    //       })
        
    // }

    handleSignIn = (props) => {
        console.log('handleSignIn called')
        .then((signInObj) => {
            if (signInObj) {
                this.setState({signedIn: true})
                console.log(signInObj.user)
            } 
        })
          .catch((err) => {
            this.setState({errorMessage: err.message});
          })

    }

    // handleSignOut = (props) => {
    //     .then( () => {
    //         this.setState({signedIn: false})
    //     })
    //     .catch((err) => {
    //         this.setState({errorMessage: err.message});
    //     })
    // }

    // addPlant = (plant) => {
    //     console.log("addPlant called")
    //     console.log(plant);
    //     console.log(this.state);
    // }

    render() {
        // console.log(this.toggleModal)

        let content = null;
        // if (!this.state.signedIn){
        //     content = (
        //         <Route
        //         path="/"
        //         render={
        //             (props) =>
        //             <LoginForm signUpCallback={this.handleSignUp} signInCallback={this.handleSignIn}/>
        //         } /> )
        // } else {
            content = (
                <div>
                    <Header addPlantCallback={this.addPlant} handleSignOutCallback={this.handleSignOut}></Header>
                    
                    <Route
                        path="/"
                        render={ 
                            (props) => 
                                <PlantList plants={this.state.plants}/>
                        } 
                    />
                </div>
            );
        //}

        return (<div>{content}</div>);
    }
}

class Header extends Component {
    constructor(props){
        super(props);
        this.addPlantCallback = this.props.addPlantCallback;
        this.handleSignOutCallback = this.props.handleSignOutCallback;
    }

    render() {
        return (
            <header>
                <nav className="navbar">
                    <span><h1 className="navbar-brand">Plant Tracker</h1></span>

                    <AddPlantModal addPlantCallback={this.addPlantCallback} toggleModal={this.toggleModal} isModalOpen={false}></AddPlantModal>
                    <SignOut handleSignOutCallback={this.handleSignOutCallback}></SignOut>
                    
                </nav>
            </header>
        );
    }
}

// class Footer extends Component {
//     render() {
//         return (
//             <footer className="page-footer">
//                 <div>
//                     <p id="footer">&copy;2021 Eric Gabrielson</p>	
//                 </div>
//             </footer>
//         );
//     }
// }

class SignOut extends Component {
    constructor(props){
        super(props);
        this.handleSignOutCallback = this.props.handleSignOutCallback;
    }
    render() {
        return (
            <Button onClick={this.handleSignOutCallback} type="button" className="btn btn-danger">
                <NavLink to="/">Sign-out</NavLink>
            </Button>
        );
    }
}

export default App;