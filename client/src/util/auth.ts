import { getUserCredentials } from "./storage";

export async function isLoginValid(): Promise<boolean> {
    const creds = getUserCredentials();
    if (!creds) return false

    const req_options = {
        method: 'GET',
    }

    let is_login_valid = false;
    await fetch(`http://localhost:8868/v1/auth/check_token?access_token=${creds.token}`, req_options)
        .then(rsp => {
            if (rsp.status === 200) {
                is_login_valid = true;
            }
        })
        .catch(e => {
            console.log(e)
        })
    
    return is_login_valid
}