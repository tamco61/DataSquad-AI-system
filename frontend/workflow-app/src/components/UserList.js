import React, { useEffect, useState } from 'react';
import { getUsers } from '../api';
import { Link } from 'react-router-dom';

const UserList = () => {
    const [users, setUsers] = useState([]);

    useEffect(() => {
        const fetchUsers = async () => {
            try {
                const response = await getUsers();
                setUsers(response.data);
            } catch (error) {
                console.error("Error fetching users:", error);
            }
        };

        fetchUsers();
    }, []);

    return (
        <div>
            <h1>Users</h1>
            <Link to="/users/new">Create User</Link>
            <ul>
                {users.map((user) => (
                    <li key={user.id}>
                        {user.name} ({user.email}) - {user.departament}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default UserList;
