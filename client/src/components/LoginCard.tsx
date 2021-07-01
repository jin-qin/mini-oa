import React from 'react';
import { Button, Form } from 'react-bootstrap';
import './LoginCard.css';

function LoginCard() {
    return(
        <div className='LoginCard'>
            <div className='LoginCard-Left'/>

            <div className='LoginCard-Right'>
                <div className='LoginCard-Top'>
                    <div className='LoginCard-Top-Title'>
                        Login
                    </div>
                    <div className='LoginCard-Top-SubTitle'>
                        Please login to continue
                    </div>
                </div>

                <div className='LoginCard-Bottom'>
                    <Form>
                        <Form.Group controlId="formBasicEmail">
                            <Form.Label style={{fontWeight: 'bold'}} className='LoginCard-Bottom-LeftText'>EMAIL</Form.Label>
                            <Form.Control type="email" placeholder="Enter email" />
                            <div className='LoginCard-Bottom-LeftText'>
                                <Form.Text className="text-muted">
                                We'll never share your email with anyone else.
                                </Form.Text>
                            </div>
                            
                        </Form.Group>

                        <Form.Group controlId="formBasicPassword">
                            <Form.Label style={{fontWeight: 'bold'}} className="LoginCard-Bottom-LeftText" >PASSWORD</Form.Label>
                            <Form.Control type="password" placeholder="Password" />
                        </Form.Group>
                        <Form.Group className="LoginCard-Bottom-Checkbox" controlId="formBasicCheckbox">
                            <Form.Check type="checkbox" label="Save password" />
                        </Form.Group>

                        <div className='LoginCard-Bottom-Btn-Submit-Layout'>
                            <Button bsPrefix="LoginCard-Bottom-Btn-Submit" variant="primary" type="submit" size='lg' >
                                LOGIN
                            </Button>
                        </div>
                    </Form>
                </div>
            </div>
        </div>
    );
}

export default LoginCard;