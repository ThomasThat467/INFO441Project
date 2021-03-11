import React, { Component } from 'react';
import api from '../../../../Constants/APIEndpoints';
import Errors from '../../../Errors/Errors';


class AddPlant extends Component () {
    constructor(props) {
        super(props);
        this.state = {
            file: null,
            error: '',
            isModalOpen: props.isModalOpen,
            plantName: '',
            wateringSchedule:[],
            img: ''
        }
    }

    sendRequest = async (e) => {
        e.preventDefault();
        const { file } = this.state;
        let data = new FormData()
        data.append('uploadfile', file);
        const response = await fetch(api.base + api.handlers.myuserPlant, {
            method: "PUT",
            body: data,
            headers: new Headers({
                "Authorization": localStorage.getItem("Authorization"),
            })
        });
        if (response.status >= 300) {
            const error = await response.text();
            console.log(error);
            this.setError(error);
            return;
        }
        alert("Plant Added"); // TODO make this better by refactoring errors
    }

    handleFile = (e) => {
        this.setState({
            file: e.target.files[0]
        })
    }

    setError = (error) => {
        this.setState({ error })
    }

     handleWateringSchedule = (schedule) => {
        console.log("handleWateringScheduleCalled", schedule);
        this.setState({wateringSchedule: schedule});
    }

     handleChange = (event) => {
        let field = event.target.name; //which input
        let value = event.target.value; //what value
    
        let changes = {}; //object to hold changes
        changes[field] = value; //change this field
        this.setState(changes); //update state
    }

//need to get this working
    toggleModal() {
        console.log("toggleModal called")
        this.setState({isModalOpen: !this.state.isModalOpen});
    }

    render() {
        const { error } = this.state;
        return <>
            <Errors error={error} setError={this.setError} />
            <div>
                <Button onClick={this.toggleModal} className="btn btn-primary">
                    Add Plant
                </Button>
                <Modal isOpen={this.state.isModalOpen} toggle={this.toggleModal}>
                    <ModalHeader>
                        Create New Plant
                        <Button onClick={this.toggleModal} type="button" className="close" aria-label="Close">
                            <span aria-hidden="true">&times;</span>
                        </Button>
                    </ModalHeader>
                    <ModalBody>
                        <form>
                            <label htmlFor="plantName">Plant Name</label>
                            <input onChange={this.handleChange} type="text" name="plantName" id="plantName" value={this.state.plantName}/>
                            <br/>
                            <div className="input-group">
                                <p>Watering Schedule</p>
                                <WateringSchedule handleWateringSchedule={this.handleWateringSchedule} modifiable={true} value={this.state.wateringSchedule}></WateringSchedule>
                            </div>

                            <div className="input-group">
                                <div className="custom-file">
                                    <label htmlFor="customFile" className="custom-file-label">Plant picture</label>
                                    <input onChange={this.handleChange} type="file" name="fileUpload" value={this.state.img} className="custom-file-input" id="customFile"/>	
                                </div>
                            </div>
                        </form>
                    </ModalBody>
                    <ModalFooter>
                        <Button onClick={this.toggleModal} className="btn btn-secondary">Close</Button>
                        <Button onClick={this.addPlant} id="newPlant" className="btn btn-primary">Save New Plant</Button>
                    </ModalFooter>
                </Modal>
            </div>
        </>
    }

}

export default AddPlant;