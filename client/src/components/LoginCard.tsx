import React, { useState } from 'react'
import { Button, Form } from 'react-bootstrap';
import './LoginCard.css';

import * as StorageUtils from '../util/storage';

function LoginCard() {
    const [username, setUsername] = useState<string>('')
    const [password, setPassword] = useState<string>('')
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
                        <Form.Group controlId="formBasicUsername">
                            <Form.Label style={{fontWeight: 'bold'}} className='LoginCard-Bottom-LeftText'>USER NAME</Form.Label>
                            <Form.Control type="text" placeholder="Enter your user name" onChange={e => setUsername(e.target.value)} />
                            <div className='LoginCard-Bottom-LeftText'>
                                <Form.Text className="text-muted">
                                We'll never share your user name with anyone else.
                                </Form.Text>
                            </div>
                        </Form.Group>

                        <Form.Group controlId="formBasicPassword">
                            <Form.Label style={{fontWeight: 'bold'}} className="LoginCard-Bottom-LeftText" >PASSWORD</Form.Label>
                            <Form.Control type="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />
                        </Form.Group>

                        <div className='LoginCard-Bottom-Btn-Submit-Layout'>
                            <Button 
                                bsPrefix="LoginCard-Bottom-Btn-Submit" 
                                variant="primary"
                                size='lg' 
                                onClick={e => onLogin(e, {username, password})}
                            >
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

interface LoginProps {
    username: string,
    password: string
}

async function onLogin(e: React.MouseEvent<HTMLElement, MouseEvent>, props: LoginProps) {
    e.preventDefault();

    const req_options = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            user: {
                username: props.username,
                password: props.password
            }
        })
    }

    fetch('http://localhost:8868/v1/users/login', req_options)
        .then(rsp => rsp.json())
        .then(data => {
            StorageUtils.saveUserCredentials(data['user']['user_id'], data['user']['username'], data['user']['token'])
        })
        .catch(e => {
            console.log(e)
        })
}