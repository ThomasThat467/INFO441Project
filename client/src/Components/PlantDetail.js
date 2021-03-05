import React, { Component, useState} from 'react';
import {Button, Modal, ModalHeader, ModalBody, ModalFooter} from 'reactstrap'

export class PlantDetailModal extends Component {
    constructor(props) {
        super(props);
        this.state = {isModalOpen: props.isModalOpen};
        // this.toggleModal = this.props.toggleModal;
        this.toggleModal = this.toggleModal.bind(this);
    }

    toggleModal() {
        this.setState({isModalOpen: !this.state.isModalOpen});
    }

    render() {

        return (
            <div>
                <Button onClick={this.toggleModal} className="btn btn-primary">
                    Add Plant
                </Button>
                <Modal isOpen={this.state.isModalOpen} toggle={this.toggleModal}>
                    <ModalHeader>
                        <h5 className="modal-title" id="plantDetailModalLongTitle">Create New Plant</h5>
                        <Button type="button" className="close" aria-label="Close">
                            <span aria-hidden="true">&times;</span>
                        </Button>
                    </ModalHeader>
                    <ModalBody>
                        <form>
                            <label htmlFor="plantName">Plant Name</label>
                            <input type="text" name="plantName" id="plantName"/>
                            <br/>
                            <div className="input-group">
                                <p>Watering Schedule</p>
                                {/* <WateringSchedule></WateringSchedule> */}
                            </div>

                            <div className="input-group">
                                <div className="custom-file">
                                    <label htmlFor="customFile" className="custom-file-label">Plant picture</label>
                                    <input type="file" name="fileUpload" className="custom-file-input" id="customFile"/>	
                                </div>
                            </div>
                        </form>
                    </ModalBody>
                    <ModalFooter>
                        <Button onClick={this.toggleModal} className="btn btn-secondary">Close</Button>
                        <Button onClick={this.toggleModal} id="newPlant" className="btn btn-primary">Save New Plant</Button>
                    </ModalFooter>
                </Modal>
            </div>
        );
    }
}