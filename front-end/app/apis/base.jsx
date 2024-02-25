

const UI_BASE_PATH = ''

const redirectLocal = () => {
    window.location = '/auth'
}


const SERVER_BASE_PATH = 'http://localhost:8099';

const AUTH_PATH = `${SERVER_BASE_PATH}/auth`;

const USER_PATH = `${SERVER_BASE_PATH}/user`;

const AUTH_TOKEN = 'token'

export {
    SERVER_BASE_PATH,
    AUTH_PATH,
    USER_PATH,
    AUTH_TOKEN
}