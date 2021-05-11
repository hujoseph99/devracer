import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useHistory } from 'react-router';

import { githubCallback, selectIsLoggedIn } from './authSlice';


export const GitHubCallback = (): JSX.Element => {
    const dispatch = useDispatch();
    const loggedin = useSelector(selectIsLoggedIn);
    const history = useHistory();

    useEffect(() => {
        dispatch(githubCallback(new URLSearchParams(window.location.search)));
    }, [dispatch])

    useEffect(() => {
        if (loggedin) {
            history.push('/')
        }
    }, [loggedin, history])
    
    return <p>Logging in</p>
}

