import { RANKS } from "$lib/constants";
import type { UserData } from "$lib/types";

export function isUserDataValid(data: UserData)
{
    return data.userId !== "" && data.username !== "";
}

export function convertTrophiesToLeague(trophies: number)
{
    const result = RANKS.filter(rank => trophies >= rank.minTrophy).pop();

    if (result === undefined)
        return null;

    return {
        league: result.league,
        division: result.division,
        stars: trophies - result.minTrophy
    };
}