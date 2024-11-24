import axios from "axios";
const BASE_URL = "http://localhost:5000";
export async function register(email, password, role) {
    try {
        const params = {
            "userEmail": email,
            "userPassword": password,
            "userRole": role
        }
        const response = await axios.post(BASE_URL + '/api/auth/register', params);
        const data = await response.data; 
        if('userId' in data){
            return login(email, password)
        }else{
            return "error";
        }

    } catch (error) {
        console.error('Ошибка при получении данных:', error.message);
        throw error; 
    }
}

export async function login(email, password) {
    try {
        const params = {
            "userEmail": email,
            "userPassword": password,
        };
        const response = await axios.post(BASE_URL + '/api/auth/login', params);
        const data = await response.data; 
        if('token' in data){
            return data; 
        }else{
            return "error";
        }
    } catch (error) {
        console.error('Ошибка при получении данных:', error.message);
        throw error;  
    }
}


export async function get_all(token) {
    try {
        const response = await axios.get(BASE_URL + '/api/get_all', {
            params: {
            "token": token,
            }
        });
        const data = await response.data.computers;
        return data; 
          
    } catch (error) {
        console.error('Ошибка при получении данных:', error.message);
        throw error;  
    }
}

export async function get_computer(id, token) {
    try {

        const response = await axios.get(BASE_URL + '/api/get_computer', {
            params :{
                "id": id,
                "token": token,
            }
        });
        const data = await response.data;
        if('ssh' in data){
            return data; 
        }
        else{
            return "error";
        }
        
          
    } catch (error) {
        console.error('Ошибка при получении данных:', error.message);
        throw error;  
    }
}

export async function reserve_computer(id, token) {
    try {
        console.log(id);
        const params = {
            "id": id,
            "token": token,
        }
        const response = await axios.post(BASE_URL + '/api/reserve_computer', params);
        const data = await response.data;
        if('reserved' in data){
            return true; 
        }
        else
        {
            return false;
        }
    } catch (error) {
        console.error('Ошибка при получении данных:', error.message);
        throw error;  
    }
}

export async function relieve_computer(id, token) {
    try {
        const params = {
            "id": id,
            "token": token,
        }
        const response = await axios.post(BASE_URL + '/api/relieve_computer', params);
        const data = await response.data;
        if('reserved' in data){
            return true; 
        }
        else
        {
            return false;
        }
    } catch (error) {
        console.error('Ошибка при получении данных:', error.message);
        throw error;  
    }
}