package lol

import "fmt"


   // "region": "NA",
   // "matchType": "MATCHED_GAME",
   // "matchCreation": 1462675467034,
   // "participants": [
   //    {
   //       "stats": {
   //          "unrealKills": 0,
   //          "item2": 1001,
   //          "item1": 3508,
   //          "totalDamageTaken": 15274,
   //          "item0": 1055,
   //          "pentaKills": 0,
   //          "sightWardsBoughtInGame": 0,
   //          "winner": false,
   //          "magicDamageDealt": 11864,
   //          "wardsKilled": 1,
   //          "largestCriticalStrike": 463,
   //          "trueDamageDealt": 0,
   //          "doubleKills": 0,
   //          "physicalDamageDealt": 108491,
   //          "tripleKills": 0,
   //          "deaths": 6,
   //          "firstBloodAssist": false,
   //          "magicDamageDealtToChampions": 826,
   //          "assists": 6,
   //          "visionWardsBoughtInGame": 0,
   //          "totalTimeCrowdControlDealt": 0,
   //          "champLevel": 13,
   //          "physicalDamageTaken": 10494,
   //          "totalDamageDealt": 120356,
   //          "largestKillingSpree": 3,
   //          "inhibitorKills": 0,
   //          "minionsKilled": 208,
   //          "towerKills": 0,
   //          "physicalDamageDealtToChampions": 9801,
   //          "quadraKills": 0,
   //          "goldSpent": 9250,
   //          "totalDamageDealtToChampions": 10627,
   //          "goldEarned": 10485,
   //          "neutralMinionsKilledTeamJungle": 7,
   //          "firstBloodKill": false,
   //          "firstTowerKill": false,
   //          "wardsPlaced": 9,
   //          "trueDamageDealtToChampions": 0,
   //          "killingSprees": 1,
   //          "firstInhibitorKill": false,
   //          "totalScoreRank": 0,
   //          "totalUnitsHealed": 3,
   //          "kills": 4,
   //          "firstInhibitorAssist": false,
   //          "totalPlayerScore": 0,
   //          "neutralMinionsKilledEnemyJungle": 0,
   //          "magicDamageTaken": 3639,
   //          "largestMultiKill": 1,
   //          "totalHeal": 2639,
   //          "item4": 1038,
   //          "item3": 3087,
   //          "objectivePlayerScore": 0,
   //          "item6": 3340,
   //          "firstTowerAssist": false,
   //          "item5": 1053,
   //          "trueDamageTaken": 1140,
   //          "neutralMinionsKilled": 7,
   //          "combatPlayerScore": 0
   //       },
   //       "timeline": {
   //          "xpDiffPerMinDeltas": {
   //             "zeroToTen": -12.700000000000003,
   //             "tenToTwenty": -11.450000000000017
   //          },
   //          "damageTakenDiffPerMinDeltas": {
   //             "zeroToTen": 29.950000000000017,
   //             "tenToTwenty": -2.0000000000000284
   //          },
   //          "xpPerMinDeltas": {
   //             "zeroToTen": 300.8,
   //             "tenToTwenty": 481.8
   //          },
   //          "goldPerMinDeltas": {
   //             "zeroToTen": 235.1,
   //             "tenToTwenty": 403.20000000000005
   //          },
   //          "role": "DUO_CARRY",
   //          "creepsPerMinDeltas": {
   //             "zeroToTen": 6.2,
   //             "tenToTwenty": 9.2
   //          },
   //          "csDiffPerMinDeltas": {
   //             "zeroToTen": -0.8,
   //             "tenToTwenty": 1.0999999999999996
   //          },
   //          "damageTakenPerMinDeltas": {
   //             "zeroToTen": 401.6,
   //             "tenToTwenty": 566.0999999999999
   //          },
   //          "lane": "BOTTOM"
   //       },
   //       "spell2Id": 7,
   //       "participantId": 1,
   //       "championId": 236,
   //       "teamId": 100,
   //       "highestAchievedSeasonTier": "SILVER",
   //       "spell1Id": 4
   //    },
   //    {
   //       "masteries": [
   //          {
   //             "rank": 5,
   //             "masteryId": 6114
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6122
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6134
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6141
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6212
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6223
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6232
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6242
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6251
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6261
   //          }
   //       ],
   //       "stats": {
   //          "unrealKills": 0,
   //          "item2": 0,
   //          "item1": 3087,
   //          "totalDamageTaken": 21077,
   //          "item0": 3508,
   //          "pentaKills": 0,
   //          "sightWardsBoughtInGame": 0,
   //          "winner": false,
   //          "magicDamageDealt": 5348,
   //          "wardsKilled": 0,
   //          "largestCriticalStrike": 448,
   //          "trueDamageDealt": 0,
   //          "doubleKills": 0,
   //          "physicalDamageDealt": 118290,
   //          "tripleKills": 0,
   //          "deaths": 5,
   //          "firstBloodAssist": false,
   //          "magicDamageDealtToChampions": 908,
   //          "assists": 4,
   //          "visionWardsBoughtInGame": 3,
   //          "totalTimeCrowdControlDealt": 89,
   //          "champLevel": 14,
   //          "physicalDamageTaken": 10120,
   //          "totalDamageDealt": 123639,
   //          "largestKillingSpree": 0,
   //          "inhibitorKills": 0,
   //          "minionsKilled": 203,
   //          "towerKills": 0,
   //          "physicalDamageDealtToChampions": 11658,
   //          "quadraKills": 0,
   //          "goldSpent": 8575,
   //          "totalDamageDealtToChampions": 12567,
   //          "goldEarned": 9734,
   //          "neutralMinionsKilledTeamJungle": 6,
   //          "firstBloodKill": false,
   //          "firstTowerKill": false,
   //          "wardsPlaced": 12,
   //          "trueDamageDealtToChampions": 0,
   //          "killingSprees": 0,
   //          "firstInhibitorKill": false,
   //          "totalScoreRank": 0,
   //          "totalUnitsHealed": 1,
   //          "kills": 1,
   //          "firstInhibitorAssist": false,
   //          "totalPlayerScore": 0,
   //          "neutralMinionsKilledEnemyJungle": 3,
   //          "magicDamageTaken": 10458,
   //          "largestMultiKill": 1,
   //          "totalHeal": 4204,
   //          "item4": 3009,
   //          "item3": 0,
   //          "objectivePlayerScore": 0,
   //          "item6": 3340,
   //          "firstTowerAssist": false,
   //          "item5": 1055,
   //          "trueDamageTaken": 498,
   //          "neutralMinionsKilled": 9,
   //          "combatPlayerScore": 0
   //       },
   //       "runes": [
   //          {
   //             "rank": 1,
   //             "runeId": 5213
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5245
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5277
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5317
   //          },
   //          {
   //             "rank": 2,
   //             "runeId": 5335
   //          }
   //       ],
   //       "timeline": {
   //          "xpDiffPerMinDeltas": {
   //             "zeroToTen": -30.50000000000003,
   //             "tenToTwenty": -109.9
   //          },
   //          "damageTakenDiffPerMinDeltas": {
   //             "zeroToTen": 33.5,
   //             "tenToTwenty": 201.40000000000003
   //          },
   //          "xpPerMinDeltas": {
   //             "zeroToTen": 419.5,
   //             "tenToTwenty": 446.20000000000005
   //          },
   //          "goldPerMinDeltas": {
   //             "zeroToTen": 254.4,
   //             "tenToTwenty": 377.20000000000005
   //          },
   //          "role": "SOLO",
   //          "creepsPerMinDeltas": {
   //             "zeroToTen": 7.8,
   //             "tenToTwenty": 7.7
   //          },
   //          "csDiffPerMinDeltas": {
   //             "zeroToTen": 2.3,
   //             "tenToTwenty": -0.2999999999999994
   //          },
   //          "damageTakenPerMinDeltas": {
   //             "zeroToTen": 358.5,
   //             "tenToTwenty": 820.5
   //          },
   //          "lane": "TOP"
   //       },
   //       "spell2Id": 12,
   //       "participantId": 2,
   //       "championId": 104,
   //       "teamId": 100,
   //       "highestAchievedSeasonTier": "GOLD",
   //       "spell1Id": 4
   //    },
   //    {
   //       "masteries": [
   //          {
   //             "rank": 5,
   //             "masteryId": 6211
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6221
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6231
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6242
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6311
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6322
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6332
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6342
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6352
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6363
   //          }
   //       ],
   //       "stats": {
   //          "unrealKills": 0,
   //          "item2": 3174,
   //          "item1": 2010,
   //          "totalDamageTaken": 10313,
   //          "item0": 2301,
   //          "pentaKills": 0,
   //          "sightWardsBoughtInGame": 0,
   //          "winner": false,
   //          "magicDamageDealt": 8857,
   //          "wardsKilled": 4,
   //          "largestCriticalStrike": 0,
   //          "trueDamageDealt": 0,
   //          "doubleKills": 0,
   //          "physicalDamageDealt": 2562,
   //          "tripleKills": 0,
   //          "deaths": 5,
   //          "firstBloodAssist": false,
   //          "magicDamageDealtToChampions": 6178,
   //          "assists": 8,
   //          "visionWardsBoughtInGame": 3,
   //          "totalTimeCrowdControlDealt": 29,
   //          "champLevel": 12,
   //          "physicalDamageTaken": 4539,
   //          "totalDamageDealt": 11420,
   //          "largestKillingSpree": 0,
   //          "inhibitorKills": 0,
   //          "minionsKilled": 10,
   //          "towerKills": 0,
   //          "physicalDamageDealtToChampions": 1548,
   //          "quadraKills": 0,
   //          "goldSpent": 6410,
   //          "totalDamageDealtToChampions": 7726,
   //          "goldEarned": 7268,
   //          "neutralMinionsKilledTeamJungle": 1,
   //          "firstBloodKill": false,
   //          "firstTowerKill": false,
   //          "wardsPlaced": 19,
   //          "trueDamageDealtToChampions": 0,
   //          "killingSprees": 0,
   //          "firstInhibitorKill": false,
   //          "totalScoreRank": 0,
   //          "totalUnitsHealed": 5,
   //          "kills": 1,
   //          "firstInhibitorAssist": false,
   //          "totalPlayerScore": 0,
   //          "neutralMinionsKilledEnemyJungle": 0,
   //          "magicDamageTaken": 5158,
   //          "largestMultiKill": 1,
   //          "totalHeal": 5191,
   //          "item4": 1052,
   //          "item3": 3020,
   //          "objectivePlayerScore": 0,
   //          "item6": 3364,
   //          "firstTowerAssist": false,
   //          "item5": 0,
   //          "trueDamageTaken": 615,
   //          "neutralMinionsKilled": 1,
   //          "combatPlayerScore": 0
   //       },
   //       "runes": [
   //          {
   //             "rank": 2,
   //             "runeId": 5167
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5273
   //          },
   //          {
   //             "rank": 7,
   //             "runeId": 5290
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5317
   //          },
   //          {
   //             "rank": 3,
   //             "runeId": 5357
   //          }
   //       ],
   //       "timeline": {
   //          "xpDiffPerMinDeltas": {
   //             "zeroToTen": -12.700000000000003,
   //             "tenToTwenty": -11.450000000000017
   //          },
   //          "damageTakenDiffPerMinDeltas": {
   //             "zeroToTen": 29.950000000000017,
   //             "tenToTwenty": -2.0000000000000284
   //          },
   //          "xpPerMinDeltas": {
   //             "zeroToTen": 262.5,
   //             "tenToTwenty": 303.1
   //          },
   //          "goldPerMinDeltas": {
   //             "zeroToTen": 162.3,
   //             "tenToTwenty": 289.4
   //          },
   //          "role": "DUO_SUPPORT",
   //          "creepsPerMinDeltas": {
   //             "zeroToTen": 0.3,
   //             "tenToTwenty": 0.30000000000000004
   //          },
   //          "csDiffPerMinDeltas": {
   //             "zeroToTen": -0.8,
   //             "tenToTwenty": 1.0999999999999996
   //          },
   //          "damageTakenPerMinDeltas": {
   //             "zeroToTen": 287.2,
   //             "tenToTwenty": 320.5
   //          },
   //          "lane": "BOTTOM"
   //       },
   //       "spell2Id": 3,
   //       "participantId": 3,
   //       "championId": 37,
   //       "teamId": 100,
   //       "highestAchievedSeasonTier": "SILVER",
   //       "spell1Id": 4
   //    },
   //    {
   //       "masteries": [
   //          {
   //             "rank": 5,
   //             "masteryId": 6212
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6221
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6231
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6241
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6251
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6262
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6311
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6321
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6331
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6343
   //          }
   //       ],
   //       "stats": {
   //          "unrealKills": 0,
   //          "item2": 3009,
   //          "item1": 2032,
   //          "totalDamageTaken": 28460,
   //          "item0": 1401,
   //          "pentaKills": 0,
   //          "sightWardsBoughtInGame": 0,
   //          "winner": false,
   //          "magicDamageDealt": 62779,
   //          "wardsKilled": 2,
   //          "largestCriticalStrike": 0,
   //          "trueDamageDealt": 13011,
   //          "doubleKills": 1,
   //          "physicalDamageDealt": 71406,
   //          "tripleKills": 0,
   //          "deaths": 2,
   //          "firstBloodAssist": false,
   //          "magicDamageDealtToChampions": 6174,
   //          "assists": 5,
   //          "visionWardsBoughtInGame": 0,
   //          "totalTimeCrowdControlDealt": 1414,
   //          "champLevel": 15,
   //          "physicalDamageTaken": 18896,
   //          "totalDamageDealt": 147196,
   //          "largestKillingSpree": 3,
   //          "inhibitorKills": 0,
   //          "minionsKilled": 68,
   //          "towerKills": 0,
   //          "physicalDamageDealtToChampions": 6216,
   //          "quadraKills": 0,
   //          "goldSpent": 8825,
   //          "totalDamageDealtToChampions": 13626,
   //          "goldEarned": 10728,
   //          "neutralMinionsKilledTeamJungle": 83,
   //          "firstBloodKill": false,
   //          "firstTowerKill": false,
   //          "wardsPlaced": 3,
   //          "trueDamageDealtToChampions": 1236,
   //          "killingSprees": 2,
   //          "firstInhibitorKill": false,
   //          "totalScoreRank": 0,
   //          "totalUnitsHealed": 1,
   //          "kills": 5,
   //          "firstInhibitorAssist": false,
   //          "totalPlayerScore": 0,
   //          "neutralMinionsKilledEnemyJungle": 3,
   //          "magicDamageTaken": 7538,
   //          "largestMultiKill": 2,
   //          "totalHeal": 9197,
   //          "item4": 3211,
   //          "item3": 3742,
   //          "objectivePlayerScore": 0,
   //          "item6": 3364,
   //          "firstTowerAssist": false,
   //          "item5": 3067,
   //          "trueDamageTaken": 2026,
   //          "neutralMinionsKilled": 86,
   //          "combatPlayerScore": 0
   //       },
   //       "runes": [
   //          {
   //             "rank": 9,
   //             "runeId": 5245
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5289
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5317
   //          },
   //          {
   //             "rank": 3,
   //             "runeId": 5337
   //          }
   //       ],
   //       "timeline": {
   //          "xpDiffPerMinDeltas": {
   //             "zeroToTen": 97.80000000000001,
   //             "tenToTwenty": 70.99999999999997
   //          },
   //          "damageTakenDiffPerMinDeltas": {
   //             "zeroToTen": -132.7,
   //             "tenToTwenty": 92.19999999999993
   //          },
   //          "xpPerMinDeltas": {
   //             "zeroToTen": 433.6,
   //             "tenToTwenty": 545.2
   //          },
   //          "goldPerMinDeltas": {
   //             "zeroToTen": 327,
   //             "tenToTwenty": 376
   //          },
   //          "role": "NONE",
   //          "creepsPerMinDeltas": {
   //             "zeroToTen": 1.7,
   //             "tenToTwenty": 2.8
   //          },
   //          "csDiffPerMinDeltas": {
   //             "zeroToTen": 1.4999999999999998,
   //             "tenToTwenty": 1.1
   //          },
   //          "damageTakenPerMinDeltas": {
   //             "zeroToTen": 546,
   //             "tenToTwenty": 1063.6999999999998
   //          },
   //          "lane": "JUNGLE"
   //       },
   //       "spell2Id": 4,
   //       "participantId": 4,
   //       "championId": 106,
   //       "teamId": 100,
   //       "highestAchievedSeasonTier": "PLATINUM",
   //       "spell1Id": 11
   //    },
   //    {
   //       "masteries": [
   //          {
   //             "rank": 5,
   //             "masteryId": 6114
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6121
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6134
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6142
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6311
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6322
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6331
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6343
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6351
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6362
   //          }
   //       ],
   //       "stats": {
   //          "unrealKills": 0,
   //          "item2": 0,
   //          "item1": 0,
   //          "totalDamageTaken": 21017,
   //          "item0": 3003,
   //          "pentaKills": 0,
   //          "sightWardsBoughtInGame": 0,
   //          "winner": false,
   //          "magicDamageDealt": 77143,
   //          "wardsKilled": 2,
   //          "largestCriticalStrike": 0,
   //          "trueDamageDealt": 0,
   //          "doubleKills": 0,
   //          "physicalDamageDealt": 11846,
   //          "tripleKills": 0,
   //          "deaths": 5,
   //          "firstBloodAssist": false,
   //          "magicDamageDealtToChampions": 5293,
   //          "assists": 4,
   //          "visionWardsBoughtInGame": 3,
   //          "totalTimeCrowdControlDealt": 72,
   //          "champLevel": 14,
   //          "physicalDamageTaken": 8647,
   //          "totalDamageDealt": 88990,
   //          "largestKillingSpree": 0,
   //          "inhibitorKills": 0,
   //          "minionsKilled": 142,
   //          "towerKills": 1,
   //          "physicalDamageDealtToChampions": 199,
   //          "quadraKills": 0,
   //          "goldSpent": 7575,
   //          "totalDamageDealtToChampions": 5493,
   //          "goldEarned": 7921,
   //          "neutralMinionsKilledTeamJungle": 0,
   //          "firstBloodKill": false,
   //          "firstTowerKill": false,
   //          "wardsPlaced": 11,
   //          "trueDamageDealtToChampions": 0,
   //          "killingSprees": 0,
   //          "firstInhibitorKill": false,
   //          "totalScoreRank": 0,
   //          "totalUnitsHealed": 1,
   //          "kills": 0,
   //          "firstInhibitorAssist": false,
   //          "totalPlayerScore": 0,
   //          "neutralMinionsKilledEnemyJungle": 0,
   //          "magicDamageTaken": 10469,
   //          "largestMultiKill": 0,
   //          "totalHeal": 6961,
   //          "item4": 3027,
   //          "item3": 1056,
   //          "objectivePlayerScore": 0,
   //          "item6": 3363,
   //          "firstTowerAssist": false,
   //          "item5": 3158,
   //          "trueDamageTaken": 1899,
   //          "neutralMinionsKilled": 0,
   //          "combatPlayerScore": 0
   //       },
   //       "runes": [
   //          {
   //             "rank": 9,
   //             "runeId": 5273
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5298
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5316
   //          },
   //          {
   //             "rank": 1,
   //             "runeId": 5357
   //          },
   //          {
   //             "rank": 2,
   //             "runeId": 5363
   //          }
   //       ],
   //       "timeline": {
   //          "xpDiffPerMinDeltas": {
   //             "zeroToTen": -122.80000000000001,
   //             "tenToTwenty": -25.600000000000023
   //          },
   //          "damageTakenDiffPerMinDeltas": {
   //             "zeroToTen": 111.30000000000001,
   //             "tenToTwenty": 196.2
   //          },
   //          "xpPerMinDeltas": {
   //             "zeroToTen": 362.5,
   //             "tenToTwenty": 441
   //          },
   //          "goldPerMinDeltas": {
   //             "zeroToTen": 189.7,
   //             "tenToTwenty": 302
   //          },
   //          "role": "SOLO",
   //          "creepsPerMinDeltas": {
   //             "zeroToTen": 4.6,
   //             "tenToTwenty": 6.199999999999999
   //          },
   //          "csDiffPerMinDeltas": {
   //             "zeroToTen": -1.0000000000000004,
   //             "tenToTwenty": -0.3999999999999999
   //          },
   //          "damageTakenPerMinDeltas": {
   //             "zeroToTen": 315,
   //             "tenToTwenty": 713.3
   //          },
   //          "lane": "MIDDLE"
   //       },
   //       "spell2Id": 1,
   //       "participantId": 5,
   //       "championId": 26,
   //       "teamId": 100,
   //       "highestAchievedSeasonTier": "PLATINUM",
   //       "spell1Id": 4
   //    },
   //    {
   //       "masteries": [
   //          {
   //             "rank": 5,
   //             "masteryId": 6211
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6223
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6231
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6241
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6311
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6322
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6332
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6342
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6352
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6363
   //          }
   //       ],
   //       "stats": {
   //          "unrealKills": 0,
   //          "item2": 2049,
   //          "item1": 3117,
   //          "totalDamageTaken": 12308,
   //          "item0": 3302,
   //          "pentaKills": 0,
   //          "sightWardsBoughtInGame": 0,
   //          "winner": true,
   //          "magicDamageDealt": 10657,
   //          "wardsKilled": 5,
   //          "largestCriticalStrike": 0,
   //          "trueDamageDealt": 4143,
   //          "doubleKills": 1,
   //          "physicalDamageDealt": 11095,
   //          "tripleKills": 0,
   //          "deaths": 1,
   //          "firstBloodAssist": false,
   //          "magicDamageDealtToChampions": 4651,
   //          "assists": 11,
   //          "visionWardsBoughtInGame": 3,
   //          "totalTimeCrowdControlDealt": 135,
   //          "champLevel": 13,
   //          "physicalDamageTaken": 7656,
   //          "totalDamageDealt": 25896,
   //          "largestKillingSpree": 4,
   //          "inhibitorKills": 0,
   //          "minionsKilled": 43,
   //          "towerKills": 0,
   //          "physicalDamageDealtToChampions": 4110,
   //          "quadraKills": 0,
   //          "goldSpent": 7925,
   //          "totalDamageDealtToChampions": 9363,
   //          "goldEarned": 9002,
   //          "neutralMinionsKilledTeamJungle": 0,
   //          "firstBloodKill": true,
   //          "firstTowerKill": false,
   //          "wardsPlaced": 15,
   //          "trueDamageDealtToChampions": 601,
   //          "killingSprees": 1,
   //          "firstInhibitorKill": false,
   //          "totalScoreRank": 0,
   //          "totalUnitsHealed": 5,
   //          "kills": 5,
   //          "firstInhibitorAssist": false,
   //          "totalPlayerScore": 0,
   //          "neutralMinionsKilledEnemyJungle": 0,
   //          "magicDamageTaken": 4551,
   //          "largestMultiKill": 2,
   //          "totalHeal": 1010,
   //          "item4": 3105,
   //          "item3": 3025,
   //          "objectivePlayerScore": 0,
   //          "item6": 3341,
   //          "firstTowerAssist": true,
   //          "item5": 3108,
   //          "trueDamageTaken": 100,
   //          "neutralMinionsKilled": 0,
   //          "combatPlayerScore": 0
   //       },
   //       "runes": [
   //          {
   //             "rank": 9,
   //             "runeId": 5245
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5289
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5317
   //          },
   //          {
   //             "rank": 3,
   //             "runeId": 5335
   //          }
   //       ],
   //       "timeline": {
   //          "xpDiffPerMinDeltas": {
   //             "zeroToTen": 12.700000000000003,
   //             "tenToTwenty": 11.450000000000017
   //          },
   //          "damageTakenDiffPerMinDeltas": {
   //             "zeroToTen": -29.950000000000017,
   //             "tenToTwenty": 2.0000000000000284
   //          },
   //          "xpPerMinDeltas": {
   //             "zeroToTen": 267.7,
   //             "tenToTwenty": 419.8
   //          },
   //          "goldPerMinDeltas": {
   //             "zeroToTen": 209.10000000000002,
   //             "tenToTwenty": 356.9
   //          },
   //          "role": "DUO_SUPPORT",
   //          "creepsPerMinDeltas": {
   //             "zeroToTen": 1.4,
   //             "tenToTwenty": 2
   //          },
   //          "csDiffPerMinDeltas": {
   //             "zeroToTen": 0.8,
   //             "tenToTwenty": -1.0999999999999996
   //          },
   //          "damageTakenPerMinDeltas": {
   //             "zeroToTen": 320.5,
   //             "tenToTwenty": 496.3
   //          },
   //          "lane": "BOTTOM"
   //       },
   //       "spell2Id": 4,
   //       "participantId": 6,
   //       "championId": 53,
   //       "teamId": 200,
   //       "highestAchievedSeasonTier": "SILVER",
   //       "spell1Id": 14
   //    },
   //    {
   //       "masteries": [
   //          {
   //             "rank": 5,
   //             "masteryId": 6114
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6121
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6134
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6142
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6151
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6164
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6212
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6221
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6232
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6241
   //          }
   //       ],
   //       "stats": {
   //          "unrealKills": 0,
   //          "item2": 3111,
   //          "item1": 2031,
   //          "totalDamageTaken": 21123,
   //          "item0": 1401,
   //          "pentaKills": 0,
   //          "sightWardsBoughtInGame": 0,
   //          "winner": true,
   //          "magicDamageDealt": 115452,
   //          "wardsKilled": 3,
   //          "largestCriticalStrike": 0,
   //          "trueDamageDealt": 15425,
   //          "doubleKills": 0,
   //          "physicalDamageDealt": 25775,
   //          "tripleKills": 0,
   //          "deaths": 3,
   //          "firstBloodAssist": false,
   //          "magicDamageDealtToChampions": 12624,
   //          "assists": 9,
   //          "visionWardsBoughtInGame": 0,
   //          "totalTimeCrowdControlDealt": 856,
   //          "champLevel": 15,
   //          "physicalDamageTaken": 17025,
   //          "totalDamageDealt": 156653,
   //          "largestKillingSpree": 3,
   //          "inhibitorKills": 0,
   //          "minionsKilled": 53,
   //          "towerKills": 1,
   //          "physicalDamageDealtToChampions": 1728,
   //          "quadraKills": 0,
   //          "goldSpent": 9525,
   //          "totalDamageDealtToChampions": 15197,
   //          "goldEarned": 11260,
   //          "neutralMinionsKilledTeamJungle": 81,
   //          "firstBloodKill": false,
   //          "firstTowerKill": false,
   //          "wardsPlaced": 12,
   //          "trueDamageDealtToChampions": 845,
   //          "killingSprees": 1,
   //          "firstInhibitorKill": false,
   //          "totalScoreRank": 0,
   //          "totalUnitsHealed": 1,
   //          "kills": 4,
   //          "firstInhibitorAssist": false,
   //          "totalPlayerScore": 0,
   //          "neutralMinionsKilledEnemyJungle": 13,
   //          "magicDamageTaken": 3982,
   //          "largestMultiKill": 1,
   //          "totalHeal": 4621,
   //          "item4": 3102,
   //          "item3": 3110,
   //          "objectivePlayerScore": 0,
   //          "item6": 3340,
   //          "firstTowerAssist": false,
   //          "item5": 1028,
   //          "trueDamageTaken": 116,
   //          "neutralMinionsKilled": 94,
   //          "combatPlayerScore": 0
   //       },
   //       "runes": [
   //          {
   //             "rank": 9,
   //             "runeId": 5273
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5298
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5317
   //          },
   //          {
   //             "rank": 3,
   //             "runeId": 5357
   //          }
   //       ],
   //       "timeline": {
   //          "xpDiffPerMinDeltas": {
   //             "zeroToTen": -97.80000000000001,
   //             "tenToTwenty": -70.99999999999997
   //          },
   //          "damageTakenDiffPerMinDeltas": {
   //             "zeroToTen": 132.7,
   //             "tenToTwenty": -92.19999999999993
   //          },
   //          "xpPerMinDeltas": {
   //             "zeroToTen": 335.8,
   //             "tenToTwenty": 474.20000000000005
   //          },
   //          "goldPerMinDeltas": {
   //             "zeroToTen": 258,
   //             "tenToTwenty": 351.4
   //          },
   //          "role": "NONE",
   //          "creepsPerMinDeltas": {
   //             "zeroToTen": 0.2,
   //             "tenToTwenty": 1.7
   //          },
   //          "csDiffPerMinDeltas": {
   //             "zeroToTen": -1.4999999999999998,
   //             "tenToTwenty": -1.1
   //          },
   //          "damageTakenPerMinDeltas": {
   //             "zeroToTen": 678.7,
   //             "tenToTwenty": 971.5
   //          },
   //          "lane": "JUNGLE"
   //       },
   //       "spell2Id": 11,
   //       "participantId": 7,
   //       "championId": 32,
   //       "teamId": 200,
   //       "highestAchievedSeasonTier": "GOLD",
   //       "spell1Id": 4
   //    },
   //    {
   //       "masteries": [
   //          {
   //             "rank": 5,
   //             "masteryId": 6111
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6122
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6131
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6142
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6211
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6223
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6232
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6241
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6251
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6261
   //          }
   //       ],
   //       "stats": {
   //          "unrealKills": 0,
   //          "item2": 3124,
   //          "item1": 3111,
   //          "totalDamageTaken": 15154,
   //          "item0": 3078,
   //          "pentaKills": 0,
   //          "sightWardsBoughtInGame": 0,
   //          "winner": true,
   //          "magicDamageDealt": 30121,
   //          "wardsKilled": 2,
   //          "largestCriticalStrike": 616,
   //          "trueDamageDealt": 337,
   //          "doubleKills": 0,
   //          "physicalDamageDealt": 96373,
   //          "tripleKills": 0,
   //          "deaths": 1,
   //          "firstBloodAssist": false,
   //          "magicDamageDealtToChampions": 4811,
   //          "assists": 2,
   //          "visionWardsBoughtInGame": 0,
   //          "totalTimeCrowdControlDealt": 55,
   //          "champLevel": 16,
   //          "physicalDamageTaken": 12725,
   //          "totalDamageDealt": 126832,
   //          "largestKillingSpree": 2,
   //          "inhibitorKills": 0,
   //          "minionsKilled": 209,
   //          "towerKills": 3,
   //          "physicalDamageDealtToChampions": 7099,
   //          "quadraKills": 0,
   //          "goldSpent": 10610,
   //          "totalDamageDealtToChampions": 11910,
   //          "goldEarned": 11150,
   //          "neutralMinionsKilledTeamJungle": 0,
   //          "firstBloodKill": false,
   //          "firstTowerKill": false,
   //          "wardsPlaced": 9,
   //          "trueDamageDealtToChampions": 0,
   //          "killingSprees": 1,
   //          "firstInhibitorKill": false,
   //          "totalScoreRank": 0,
   //          "totalUnitsHealed": 1,
   //          "kills": 3,
   //          "firstInhibitorAssist": false,
   //          "totalPlayerScore": 0,
   //          "neutralMinionsKilledEnemyJungle": 2,
   //          "magicDamageTaken": 1921,
   //          "largestMultiKill": 1,
   //          "totalHeal": 1865,
   //          "item4": 1028,
   //          "item3": 3082,
   //          "objectivePlayerScore": 0,
   //          "item6": 3340,
   //          "firstTowerAssist": false,
   //          "item5": 0,
   //          "trueDamageTaken": 508,
   //          "neutralMinionsKilled": 2,
   //          "combatPlayerScore": 0
   //       },
   //       "runes": [
   //          {
   //             "rank": 9,
   //             "runeId": 5245
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5290
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5317
   //          },
   //          {
   //             "rank": 3,
   //             "runeId": 5337
   //          }
   //       ],
   //       "timeline": {
   //          "xpDiffPerMinDeltas": {
   //             "zeroToTen": 30.50000000000003,
   //             "tenToTwenty": 109.9
   //          },
   //          "damageTakenDiffPerMinDeltas": {
   //             "zeroToTen": -33.5,
   //             "tenToTwenty": -201.40000000000003
   //          },
   //          "xpPerMinDeltas": {
   //             "zeroToTen": 450,
   //             "tenToTwenty": 556.1
   //          },
   //          "goldPerMinDeltas": {
   //             "zeroToTen": 212.2,
   //             "tenToTwenty": 377.79999999999995
   //          },
   //          "role": "SOLO",
   //          "creepsPerMinDeltas": {
   //             "zeroToTen": 5.5,
   //             "tenToTwenty": 8
   //          },
   //          "csDiffPerMinDeltas": {
   //             "zeroToTen": -2.3,
   //             "tenToTwenty": 0.2999999999999994
   //          },
   //          "damageTakenPerMinDeltas": {
   //             "zeroToTen": 325,
   //             "tenToTwenty": 619.0999999999999
   //          },
   //          "lane": "TOP"
   //       },
   //       "spell2Id": 12,
   //       "participantId": 8,
   //       "championId": 24,
   //       "teamId": 200,
   //       "highestAchievedSeasonTier": "SILVER",
   //       "spell1Id": 4
   //    },
   //    {
   //       "masteries": [
   //          {
   //             "rank": 5,
   //             "masteryId": 6111
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6122
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6131
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6141
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6151
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6162
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6312
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6322
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6331
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6343
   //          }
   //       ],
   //       "stats": {
   //          "unrealKills": 0,
   //          "item2": 3508,
   //          "item1": 1036,
   //          "totalDamageTaken": 13720,
   //          "item0": 3072,
   //          "pentaKills": 0,
   //          "sightWardsBoughtInGame": 0,
   //          "winner": true,
   //          "magicDamageDealt": 0,
   //          "wardsKilled": 2,
   //          "largestCriticalStrike": 442,
   //          "trueDamageDealt": 1196,
   //          "doubleKills": 0,
   //          "physicalDamageDealt": 148074,
   //          "tripleKills": 0,
   //          "deaths": 3,
   //          "firstBloodAssist": false,
   //          "magicDamageDealtToChampions": 0,
   //          "assists": 10,
   //          "visionWardsBoughtInGame": 2,
   //          "totalTimeCrowdControlDealt": 79,
   //          "champLevel": 14,
   //          "physicalDamageTaken": 10204,
   //          "totalDamageDealt": 149270,
   //          "largestKillingSpree": 3,
   //          "inhibitorKills": 0,
   //          "minionsKilled": 189,
   //          "towerKills": 0,
   //          "physicalDamageDealtToChampions": 16843,
   //          "quadraKills": 0,
   //          "goldSpent": 11850,
   //          "totalDamageDealtToChampions": 17423,
   //          "goldEarned": 11752,
   //          "neutralMinionsKilledTeamJungle": 15,
   //          "firstBloodKill": false,
   //          "firstTowerKill": false,
   //          "wardsPlaced": 13,
   //          "trueDamageDealtToChampions": 580,
   //          "killingSprees": 1,
   //          "firstInhibitorKill": false,
   //          "totalScoreRank": 0,
   //          "totalUnitsHealed": 3,
   //          "kills": 5,
   //          "firstInhibitorAssist": false,
   //          "totalPlayerScore": 0,
   //          "neutralMinionsKilledEnemyJungle": 3,
   //          "magicDamageTaken": 3135,
   //          "largestMultiKill": 1,
   //          "totalHeal": 2669,
   //          "item4": 3046,
   //          "item3": 3158,
   //          "objectivePlayerScore": 0,
   //          "item6": 3363,
   //          "firstTowerAssist": false,
   //          "item5": 0,
   //          "trueDamageTaken": 380,
   //          "neutralMinionsKilled": 18,
   //          "combatPlayerScore": 0
   //       },
   //       "runes": [
   //          {
   //             "rank": 9,
   //             "runeId": 5245
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5289
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5317
   //          },
   //          {
   //             "rank": 3,
   //             "runeId": 5337
   //          }
   //       ],
   //       "timeline": {
   //          "xpDiffPerMinDeltas": {
   //             "zeroToTen": 12.700000000000003,
   //             "tenToTwenty": 11.450000000000017
   //          },
   //          "damageTakenDiffPerMinDeltas": {
   //             "zeroToTen": -29.950000000000017,
   //             "tenToTwenty": 2.0000000000000284
   //          },
   //          "xpPerMinDeltas": {
   //             "zeroToTen": 321,
   //             "tenToTwenty": 388
   //          },
   //          "goldPerMinDeltas": {
   //             "zeroToTen": 294.4,
   //             "tenToTwenty": 389.29999999999995
   //          },
   //          "role": "DUO_CARRY",
   //          "creepsPerMinDeltas": {
   //             "zeroToTen": 6.7,
   //             "tenToTwenty": 5.3
   //          },
   //          "csDiffPerMinDeltas": {
   //             "zeroToTen": 0.8,
   //             "tenToTwenty": -1.0999999999999996
   //          },
   //          "damageTakenPerMinDeltas": {
   //             "zeroToTen": 308.4,
   //             "tenToTwenty": 394.3
   //          },
   //          "lane": "BOTTOM"
   //       },
   //       "spell2Id": 4,
   //       "participantId": 9,
   //       "championId": 15,
   //       "teamId": 200,
   //       "highestAchievedSeasonTier": "GOLD",
   //       "spell1Id": 7
   //    },
   //    {
   //       "masteries": [
   //          {
   //             "rank": 5,
   //             "masteryId": 6114
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6122
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6131
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6142
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6311
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6322
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6331
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6343
   //          },
   //          {
   //             "rank": 5,
   //             "masteryId": 6351
   //          },
   //          {
   //             "rank": 1,
   //             "masteryId": 6362
   //          }
   //       ],
   //       "stats": {
   //          "unrealKills": 0,
   //          "item2": 3020,
   //          "item1": 3165,
   //          "totalDamageTaken": 12900,
   //          "item0": 1056,
   //          "pentaKills": 0,
   //          "sightWardsBoughtInGame": 0,
   //          "winner": true,
   //          "magicDamageDealt": 69823,
   //          "wardsKilled": 1,
   //          "largestCriticalStrike": 0,
   //          "trueDamageDealt": 35575,
   //          "doubleKills": 1,
   //          "physicalDamageDealt": 17069,
   //          "tripleKills": 0,
   //          "deaths": 3,
   //          "firstBloodAssist": false,
   //          "magicDamageDealtToChampions": 15177,
   //          "assists": 6,
   //          "visionWardsBoughtInGame": 2,
   //          "totalTimeCrowdControlDealt": 47,
   //          "champLevel": 15,
   //          "physicalDamageTaken": 6978,
   //          "totalDamageDealt": 122468,
   //          "largestKillingSpree": 4,
   //          "inhibitorKills": 0,
   //          "minionsKilled": 180,
   //          "towerKills": 1,
   //          "physicalDamageDealtToChampions": 1087,
   //          "quadraKills": 0,
   //          "goldSpent": 9250,
   //          "totalDamageDealtToChampions": 20418,
   //          "goldEarned": 10705,
   //          "neutralMinionsKilledTeamJungle": 1,
   //          "firstBloodKill": false,
   //          "firstTowerKill": true,
   //          "wardsPlaced": 8,
   //          "trueDamageDealtToChampions": 4153,
   //          "killingSprees": 2,
   //          "firstInhibitorKill": false,
   //          "totalScoreRank": 0,
   //          "totalUnitsHealed": 1,
   //          "kills": 6,
   //          "firstInhibitorAssist": false,
   //          "totalPlayerScore": 0,
   //          "neutralMinionsKilledEnemyJungle": 0,
   //          "magicDamageTaken": 5789,
   //          "largestMultiKill": 2,
   //          "totalHeal": 2574,
   //          "item4": 1058,
   //          "item3": 3285,
   //          "objectivePlayerScore": 0,
   //          "item6": 3363,
   //          "firstTowerAssist": false,
   //          "item5": 2043,
   //          "trueDamageTaken": 132,
   //          "neutralMinionsKilled": 1,
   //          "combatPlayerScore": 0
   //       },
   //       "runes": [
   //          {
   //             "rank": 9,
   //             "runeId": 5273
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5289
   //          },
   //          {
   //             "rank": 9,
   //             "runeId": 5317
   //          },
   //          {
   //             "rank": 3,
   //             "runeId": 5357
   //          }
   //       ],
   //       "timeline": {
   //          "xpDiffPerMinDeltas": {
   //             "zeroToTen": 122.80000000000001,
   //             "tenToTwenty": 25.600000000000023
   //          },
   //          "damageTakenDiffPerMinDeltas": {
   //             "zeroToTen": -111.30000000000001,
   //             "tenToTwenty": -196.2
   //          },
   //          "xpPerMinDeltas": {
   //             "zeroToTen": 485.3,
   //             "tenToTwenty": 466.6
   //          },
   //          "goldPerMinDeltas": {
   //             "zeroToTen": 234.89999999999998,
   //             "tenToTwenty": 404.6
   //          },
   //          "role": "SOLO",
   //          "creepsPerMinDeltas": {
   //             "zeroToTen": 5.6,
   //             "tenToTwenty": 6.6
   //          },
   //          "csDiffPerMinDeltas": {
   //             "zeroToTen": 1.0000000000000004,
   //             "tenToTwenty": 0.3999999999999999
   //          },
   //          "damageTakenPerMinDeltas": {
   //             "zeroToTen": 203.7,
   //             "tenToTwenty": 517.1
   //          },
   //          "lane": "MIDDLE"
   //       },
   //       "spell2Id": 14,
   //       "participantId": 10,
   //       "championId": 103,
   //       "teamId": 200,
   //       "highestAchievedSeasonTier": "GOLD",
   //       "spell1Id": 4
   //    }
   // ],
   // "platformId": "NA1",
   // "matchMode": "CLASSIC",
   // "participantIdentities": [
   //    {
   //       "player": {
   //          "profileIcon": 761,
   //          "matchHistoryUri": "/v1/stats/player_history/NA1/207660566",
   //          "summonerName": "Robloxin",
   //          "summonerId": 44803803
   //       },
   //       "participantId": 1
   //    },
   //    {
   //       "player": {
   //          "profileIcon": 1117,
   //          "matchHistoryUri": "/v1/stats/player_history/NA1/208665564",
   //          "summonerName": "Berzerking",
   //          "summonerId": 45631137
   //       },
   //       "participantId": 2
   //    },
   //    {
   //       "player": {
   //          "profileIcon": 1117,
   //          "matchHistoryUri": "/v1/stats/player_history/NA1/49367243",
   //          "summonerName": "Terrorgod",
   //          "summonerId": 34833527
   //       },
   //       "participantId": 3
   //    },
   //    {
   //       "player": {
   //          "profileIcon": 1117,
   //          "matchHistoryUri": "/v1/stats/player_history/NA/40487889",
   //          "summonerName": "wooooooooooooons",
   //          "summonerId": 25924430
   //       },
   //       "participantId": 4
   //    },
   //    {
   //       "player": {
   //          "profileIcon": 544,
   //          "matchHistoryUri": "/v1/stats/player_history/NA/39518098",
   //          "summonerName": "Mrjukes",
   //          "summonerId": 25015942
   //       },
   //       "participantId": 5
   //    },
   //    {
   //       "player": {
   //          "profileIcon": 28,
   //          "matchHistoryUri": "/v1/stats/player_history/NA1/213725679",
   //          "summonerName": "Doug Judy",
   //          "summonerId": 51165505
   //       },
   //       "participantId": 6
   //    },
   //    {
   //       "player": {
   //          "profileIcon": 607,
   //          "matchHistoryUri": "/v1/stats/player_history/NA/33053666",
   //          "summonerName": "Retro Spin Kidd",
   //          "summonerId": 20189772
   //       },
   //       "participantId": 7
   //    },
   //    {
   //       "player": {
   //          "profileIcon": 1117,
   //          "matchHistoryUri": "/v1/stats/player_history/NA/41375996",
   //          "summonerName": "yaw hide",
   //          "summonerId": 26691960
   //       },
   //       "participantId": 8
   //    },
   //    {
   //       "player": {
   //          "profileIcon": 1117,
   //          "matchHistoryUri": "/v1/stats/player_history/NA1/46525497",
   //          "summonerName": "Puddle42",
   //          "summonerId": 31726077
   //       },
   //       "participantId": 9
   //    },
   //    {
   //       "player": {
   //          "profileIcon": 1117,
   //          "matchHistoryUri": "/v1/stats/player_history/NA1/204486785",
   //          "summonerName": "EuglossaCognata",
   //          "summonerId": 41711438
   //       },
   //       "participantId": 10
   //    }
   // ],
   // "matchVersion": "6.9.142.993",
   // "teams": [
   //    {
   //       "bans": [
   //          {
   //             "pickTurn": 1,
   //             "championId": 54
   //          },
   //          {
   //             "pickTurn": 3,
   //             "championId": 238
   //          },
   //          {
   //             "pickTurn": 5,
   //             "championId": 16
   //          }
   //       ],
   //       "firstBlood": false,
   //       "firstTower": false,
   //       "firstInhibitor": false,
   //       "winner": false,
   //       "firstDragon": true,
   //       "vilemawKills": 0,
   //       "baronKills": 0,
   //       "teamId": 100,
   //       "inhibitorKills": 0,
   //       "dominionVictoryScore": 0,
   //       "riftHeraldKills": 0,
   //       "firstRiftHerald": false,
   //       "towerKills": 3,
   //       "firstBaron": false,
   //       "dragonKills": 1
   //    },
   //    {
   //       "bans": [
   //          {
   //             "pickTurn": 2,
   //             "championId": 45
   //          },
   //          {
   //             "pickTurn": 4,
   //             "championId": 90
   //          },
   //          {
   //             "pickTurn": 6,
   //             "championId": 105
   //          }
   //       ],
   //       "firstBlood": true,
   //       "firstTower": true,
   //       "firstInhibitor": false,
   //       "winner": true,
   //       "firstDragon": false,
   //       "vilemawKills": 0,
   //       "baronKills": 0,
   //       "teamId": 200,
   //       "inhibitorKills": 0,
   //       "dominionVictoryScore": 0,
   //       "riftHeraldKills": 0,
   //       "firstRiftHerald": false,
   //       "towerKills": 5,
   //       "firstBaron": false,
   //       "dragonKills": 1
   //    }
   // ],
   // "mapId": 11,
   // "season": "SEASON2016",
   // "queueType": "TEAM_BUILDER_DRAFT_RANKED_5x5",
   // "matchId": 2181069408,
   // "matchDuration": 1700
type Match struct {
    Participants []struct {
        Stat struct {
            Winner bool `json:"winner"`
        }   `json:"stats"`
        Timeline struct {
            Role string `json:"role"`
            Lane string `json:"lane"`
        }   `json:"timeline"`
        ParticipantId int `json:"participantId"`
        ChampionId    int `json:"championId"`
        TeamId        int `json:"teamId"`
    }   `json:"participants"`
    ParticipantIdentities []struct {
        Player struct {
            SummonerId uint64 `json:"summonerId"`
        }
        ParticipantId int `json:"participantId"`
    }   `json:"participantIdentities"`
}

func (a *APIEndpoint) formatMatchURL(summonerId uint64, options map[string]string) string {
    res := fmt.Sprintf("https://%s/api/lol/%s/v2.2/match/%v?api_key=%s", a.region.url, a.region.code, summonerId, a.key)
    for k, v := range options {
        res = fmt.Sprintf("%s&%s=%s", res, k, v)
    }
    return res
}

func (a *APIEndpoint) GetMatch(matchId uint64, includeTimeline bool) (*Match, error) {
    res := &Match{}
    options := map[string]string{}
    options["includeTimeline"] = fmt.Sprintf("%t", includeTimeline)
    url := a.formatMatchURL(matchId, options)
    err := a.g.Get(url, &res)
    if err != nil {
        return nil, err
    }
    return res, nil
}
