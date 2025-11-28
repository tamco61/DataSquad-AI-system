import axios from 'axios';

const API_URL = 'http://localhost:5000/workflow';

export const createUser = (userData) => {
    return axios.post(`${API_URL}/users`, userData);
};

export const getUsers = () => {
    return axios.get(`${API_URL}/users`);
};

export const createTask = (taskData) => {
    return axios.post(`${API_URL}/tasks`, taskData);
};

export const getTasks = (filters = {}) => {
    return axios.get(`${API_URL}/tasks`, { params: filters });
};

export const getTaskById = (id) => {
    return axios.get(`${API_URL}/tasks/${id}`);
};

export const updateTask = (id, updateData) => {
    return axios.patch(`${API_URL}/tasks/${id}`, updateData);
};
