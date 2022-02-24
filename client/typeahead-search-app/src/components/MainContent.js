
import {Container} from 'react-bootstrap';
import { useState } from 'react';
import fetchSuggestions from '../services/httpService';
import Autosuggest from 'react-autosuggest';

function MainContent () {
    const [value, setValue] = useState('');
    const [suggestedValue, setSuggestedValue] = useState('');

    const onChange = (event, { newValue }) => {
        setValue(newValue);
    };

    // You already implemented this logic above, so just use it.
    const onSuggestionsFetchRequested = async ({ value }) => {
        let returnData = await fetchSuggestions(value);
        setSuggestedValue(returnData.data.matches);
    };

    // Autosuggest will call this function every time you need to clear suggestions.
    const onSuggestionsClearRequested = () =>
    {
        setSuggestedValue('');  
    };

    const getSuggestionValue = suggestion => suggestion;

    // Use your imagination to render suggestions.
    const renderSuggestion = function (suggestion) {
        return (
            <div>
            {suggestion}
        </div>
        );
    }
   

    // Autosuggest will pass through all these props to the input.
    const inputProps = {
        placeholder: 'Type your query',
        value,
        onChange: onChange
    };

    return (
        <>
        <Container>
            <Autosuggest
                suggestions={suggestedValue}
                onSuggestionsFetchRequested={onSuggestionsFetchRequested}
                onSuggestionsClearRequested={onSuggestionsClearRequested}
                getSuggestionValue={getSuggestionValue}
                renderSuggestion={renderSuggestion}
                inputProps={inputProps}
            />
        </Container>
        </>
    )
}

export default MainContent;