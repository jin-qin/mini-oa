import { useEffect, useState } from "react";
import { Redirect } from "react-router-dom";
import * as AuthUtils from '../util/auth';

function Dashboard() {
    const [isLoginValid, setIsLoginValid] = useState<boolean>(true)
    
    useEffect(() => {
        const login_check = AuthUtils.isLoginValid()
        login_check.then(valid => {
            if (!valid) {
                setIsLoginValid(false)
            }
        })
    }, [isLoginValid])

    return (
        !isLoginValid ? 
        <Redirect to='/login'/> 
        :
        (
            <div>
                DASHBOARD
            </div>
        )
    );
}

export default Dashboard;