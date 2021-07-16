import React, { useState } from 'react';
import { Button, Form } from 'react-bootstrap';
import './RegisterCard.css';

function RegisterCard() {
    const [username, setUsername] = useState<string>('')
    const [email, setEmail] = useState<string>('')
    const [password, setPassword] = useState<string>('')

    return(
        <div className='RegisterCard'>
            <div className='RegisterCard-Left'/>

            <div className='RegisterCard-Right'>
                <div className='RegisterCard-Top'>
                    <div className='RegisterCard-Top-Title'>
                        Register
                    </div>
                    <div className='RegisterCard-Top-SubTitle'>
                        Create a new account
                    </div>
                </div>

                <div className='RegisterCard-Bottom'>
                    <Form>
                        <Form.Group controlId="formUsername">
                            <Form.Label style={{fontWeight: 'bold'}} className='RegisterCard-Bottom-LeftText'>USERNAME</Form.Label>
                            <Form.Control type="text" placeholder="Enter user name" onChange={e => setUsername(e.target.value)}/>
                        </Form.Group>
                        <Form.Group controlId="formBasicEmail">
                            <Form.Label style={{fontWeight: 'bold'}} className='RegisterCard-Bottom-LeftText'>EMAIL</Form.Label>
                            <Form.Control type="email" placeholder="Enter email" onChange={e => setEmail(e.target.value)} />
                            <div className='RegisterCard-Bottom-LeftText'>
                                <Form.Text className="text-muted">
                                We'll never share your email with anyone else.
                                </Form.Text>
                            </div>
                        </Form.Group>

                        <Form.Group controlId="formBasicPassword">
                            <Form.Label style={{fontWeight: 'bold'}} className="RegisterCard-Bottom-LeftText" >PASSWORD</Form.Label>
                            <Form.Control type="password" placeholder="Password" onChange={e => setPassword(e.target.value)}/>
                        </Form.Group>

                        <div className='RegisterCard-Bottom-Btn-Submit-Layout'>
                            <Button 
                                bsPrefix="RegisterCard-Bottom-Btn-Submit" 
                                variant="primary" 
                                type="submit" 
                                size='lg' 
                                onClick={e => onRegister(e, {username, email, password})}
                            >
                                REGISTER
                            </Button>
                        </div>
                    </Form>
                </div>
            </div>
        </div>
    );
}

export default RegisterCard;

interface RegisterProps {
    username: string,
    email: string,
    password: string
}

async function onRegister(e: React.MouseEvent<HTMLElement, MouseEvent>, props: RegisterProps) {
    e.preventDefault();

    const req_options = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(
            {
                user: {
                    username: props.username,
                    email: props.email,
                    password: props.password
                }
            }
        )
    }

    fetch('http://localhost:8868/v1/users/register', req_options)
        .then(rsp => rsp.json())
        .then(data => {
            console.log(data)
        })
        .catch(e => {
            console.log(e)
        })
}