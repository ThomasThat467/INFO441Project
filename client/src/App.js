import React, { Component } from 'react';
import Auth from './Components/Auth/Auth';
import PageTypes from './Constants/PageTypes';
import Main from './Components/Main/Main';
import './Styles/App.css';
import api from './Constants/APIEndpoints';

class App extends Component {
    constructor() {
        super();
        this.state = {
            page: localStorage.getItem("Authorization") ? PageTypes.signedInMain : PageTypes.signIn,
            authToken: localStorage.getItem("Authorization") || null,
            user: null,
            plants: []
        }

        this.getCurrentUser()
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


    /**
     * @description Gets the users
     */
    getCurrentUser = async () => {
        if (!this.state.authToken) {
            return;
        }
        const response = await fetch(api.base + api.handlers.myuser, {
            headers: new Headers({
                "Authorization": this.state.authToken
            })
        });
        if (response.status >= 300) {
            alert("Unable to verify login. Logging out...");
            localStorage.setItem("Authorization", "");
            this.setAuthToken("");
            this.setUser(null)
            return;
        }
        const user = await response.json()
        this.setUser(user);

    }

    /**
     * @description Gets the plants
     */
         getCurrentPlants = async () => {
            if (!this.state.authToken) {
                return;
            }
            const response = await fetch(api.base + api.handlers.myplants, {
                headers: new Headers({
                    "Authorization": this.state.authToken
                })
            });
            if (response.status >= 300) {
                alert("Unable to retrieve plants...");
                // localStorage.setItem("Authorization", "");
                // this.setAuthToken("");
                // this.setUser(null)
                return;
            }
            const plants = await response.json()
            this.setPlants(plants);
    
        }

    /**
     * @description sets the page type to sign in
     */
    setPageToSignIn = (e) => {
        e.preventDefault();
        this.setState({ page: PageTypes.signIn });
    }

    /**
     * @description sets the page type to sign up
     */
    setPageToSignUp = (e) => {
        e.preventDefault();
        this.setState({ page: PageTypes.signUp });
    }

    setPage = (e, page) => {
        e.preventDefault();
        this.setState({ page });
    }

    /**
     * @description sets auth token
     */
    setAuthToken = (authToken) => {
        this.setState({ authToken, page: authToken === "" ? PageTypes.signIn : PageTypes.signedInMain });
    }

    /**
     * @description sets the user
     */
    setUser = (user) => {
        this.setState({ user });
    }

    /**
     * @description sets the plants
     */
     setPlants = (plants) => {
        this.setState({ plants });
    }

    render() {
        const { page, user, plants } = this.state;
        return (
            <div>
                {user ?
                    <Main page={page}
                        setPage={this.setPage}
                        setAuthToken={this.setAuthToken}
                        user={user}
                        setUser={this.setUser}
                        plants={plants}
                        setPlants={this.setPlants} />
                    :
                    <Auth page={page}
                        setPage={this.setPage}
                        setAuthToken={this.setAuthToken}
                        setUser={this.setUser} />
                }
            </div>
        );
    }
}

export default App;