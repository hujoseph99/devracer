import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useHistory } from 'react-router';

import { githubCallback, selectIsLoggedIn } from './authSlice';


export const GitHubCallback = (): JSX.Element => {
    const dispatch = useDispatch();
    const loggedIn = useSelector(selectIsLoggedIn);
    const history = useHistory();

    useEffect(() => {
        dispatch(githubCallback(new URLSearchParams(window.location.search)));
    }, [dispatch])

    useEffect(() => {
        if (loggedIn) {
            // remove query string from url (oauth state and code)
            window.history.replaceState(null, "", window.location.href.split("?")[0])
            history.replace('/')
        }
    }, [loggedIn, history])

    return <></>
}

