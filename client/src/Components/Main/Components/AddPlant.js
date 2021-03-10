import React, {Component} from 'react';
import {Button, Modal, ModalHeader, ModalBody, ModalFooter} from 'reactstrap'
import {WateringSchedule} from './WateringSchedule.js'

export class AddPlantModal extends Component {
    constructor(props) {
        super(props);
        this.state = {
            isModalOpen: props.isModalOpen,
            plantName: '',
            wateringSchedule:[],
            img: ''
        };
        this.toggleModal = this.toggleModal.bind(this);
        this.addPlantCallback = this.props.addPlantCallback;
        this.addPlant = this.addPlant.bind(this);
    }

    toggleModal() {
        console.log("toggleModal called")
        this.setState({isModalOpen: !this.state.isModalOpen});
    }

    // plantRefs = call to database

    addPlant() {
        let newPlant = {plantName: this.state.plantName, wateringSchedule: this.state.wateringSchedule, lastWatered: '', img: this.state.img}
        // plantsRef.push(newPlant);

        // plantsRef.on('value', (snapshot) => {
        //     let plants = snapshot.val();
        //     this.setState({plants: plants})
        // });
        this.toggleModal();
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

    render() {

        return (
            <div>
                <Button onClick={this.toggleModal} className="btn btn-primary">
                    Add Plant
                </Button>
                <Modal isOpen={this.state.isModalOpen} toggle={this.toggleModal} className="add-plant">
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
        );
    }
}