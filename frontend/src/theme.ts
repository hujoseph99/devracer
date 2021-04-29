import { createMuiTheme, responsiveFontSizes } from "@material-ui/core";
import { grey } from "@material-ui/core/colors";

export const theme = responsiveFontSizes(createMuiTheme({
	typography: {
		h3: {
			fontFamily: [
				'"Proxima Nova"',
				'"Helvetica"',
				'-apple-system',
				'Arial',
				'sans-serif',
			].join(','),
		},
		fontFamily: [
			'"Helvetica Neue"',
			'-apple-system',
			'Arial',
			'sans-serif',
		].join(','),
	},
	palette: {
		primary: {
			main: grey[50]
		},
		secondary: {
			main: grey[900]
		},
		background: {
			default: grey[900]
		},
	}
}));
