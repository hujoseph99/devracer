import { createMuiTheme } from "@material-ui/core";
import { grey } from "@material-ui/core/colors";

export const theme = createMuiTheme({
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
});
