export function saveUserCredentials(user_id: string, username: string, token: string) {
    const local_storage = window.localStorage
    local_storage.setItem('user_id', user_id)
    local_storage.setItem('username', username)
    local_storage.setItem('token', token)
}

interface UserCredentials {
    user_id: string,
    username: string,
    token: string
}
export function getUserCredentials(): UserCredentials | null {
    const local_storage = window.localStorage
    const creds : UserCredentials = {user_id: '', username: '', token: ''};

    
    const user_id = local_storage.getItem('user_id');
    const username = local_storage.getItem('username');
    const token = local_storage.getItem('token');

    if (!user_id || !username || !token) {
        return null;
    }

    creds.user_id = user_id;
    creds.username = username;
    creds.token = token;

    return creds;
}