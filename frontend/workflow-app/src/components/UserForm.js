import React, { useState } from 'react';
import { createUser } from '../api';
import { useNavigate } from 'react-router-dom';

const UserForm = () => {
    const [email, setEmail] = useState('');
    const [name, setName] = useState('');
    const [departament, setDepartament] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            await createUser({ email, name, departament });
            navigate('/users');
        } catch (error) {
            console.error('Error creating user:', error);
        }
    };

    return (
        <div>
            <h1>Create User</h1>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Email:</label>
                    <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
                </div>
                <div>
                    <label>Name:</label>
                    <input type="text" value={name} onChange={(e) => setName(e.target.value)} required />
                </div>
                <div>
                    <label>Department:</label>
                    <input type="text" value={departament} onChange={(e) => setDepartament(e.target.value)} />
                </div>
                <button type="submit">Create</button>
            </form>
        </div>
    );
};

export default UserForm;
