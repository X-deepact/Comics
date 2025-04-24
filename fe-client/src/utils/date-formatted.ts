import { formatDistanceToNow, getYear } from "date-fns";

export const getTimeAgo = (
	date: Date | string | number,
	addSuffix = true
): string => {
	try {
		const parsedDate =
			typeof date === 'string' || typeof date === 'number'
				? new Date(date)
				: date;
		
		return formatDistanceToNow(parsedDate, { addSuffix });
	} catch (error) {
		console.error('Invalid date provided to getTimeAgo:', error);
		return 'Invalid date';
	}
};

export const formatToOnlyYear = (
	date: Date | string | number,
): number => {
	return getYear(date);
}
