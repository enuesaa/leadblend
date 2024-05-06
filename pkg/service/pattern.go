package service

import (
	"encoding/json"

	"github.com/enuesaa/leadblend/pkg/repository"
	"github.com/enuesaa/leadblend/pkg/repository/dbq"
	"github.com/oklog/ulid/v2"
)

type PatternTrait struct {
	Pattern dbq.Pattern
	Traits  []dbq.Trait
}

func NewPatternService(repos repository.Repos) PatternService {
	return PatternService{
		repos: repos,
	}
}

type PatternService struct {
	repos repository.Repos
}

func (srv *PatternService) List() ([]dbq.Pattern, error) {
	query := srv.repos.DB.Query()
	return query.ListPatterns(ctx())
}

func (srv *PatternService) ListWithTrait() ([]PatternTrait, error) {
	query := srv.repos.DB.Query()
	patterns, err := query.ListPatterns(ctx())
	if err != nil {
		return make([]PatternTrait, 0), err
	}
	patternIds := make([]string, 0)
	for _, p := range patterns {
		patternIds = append(patternIds, p.ID)
	}
	traits, err := query.ListTraitsByPatternIds(ctx(), patternIds)
	if err != nil {
		return make([]PatternTrait, 0), err
	}
	list := make([]PatternTrait, 0)
	for _, p := range patterns {
		pt := PatternTrait{
			Pattern: p,
			Traits:  make([]dbq.Trait, 0),
		}
		for _, t := range traits {
			if t.PatternID == p.ID {
				pt.Traits = append(pt.Traits, t)
			}
		}
		list = append(list, pt)
	}
	return list, nil
}

func (srv *PatternService) Get(id string) (dbq.Pattern, error) {
	query := srv.repos.DB.Query()
	return query.GetPattern(ctx(), id)
}

func (srv *PatternService) GetWithTrait(id string) (PatternTrait, error) {
	query := srv.repos.DB.Query()
	pattern, err := query.GetPattern(ctx(), id)
	if err != nil {
		return PatternTrait{}, err
	}
	traits, err := query.ListTraitsByPatternIds(ctx(), []string{pattern.ID})
	if err != nil {
		return PatternTrait{}, err
	}
	pt := PatternTrait{
		Pattern: pattern,
		Traits:  traits,
	}
	return pt, nil
}

func (srv *PatternService) Create(params dbq.CreatePatternParams) (string, error) {
	query := srv.repos.DB.Query()
	params.ID = ulid.Make().String()
	if _, err := query.CreatePattern(ctx(), params); err != nil {
		return "", nil
	}
	return params.ID, nil
}

func (srv *PatternService) AddTrait(params dbq.CreateTraitParams) (string, error) {
	query := srv.repos.DB.Query()
	params.ID = ulid.Make().String()
	if _, err := query.CreateTrait(ctx(), params); err != nil {
		return "", nil
	}
	return params.ID, nil
}

func (srv *PatternService) Evaluate(stoneId string) error {
	query := srv.repos.DB.Query()
	stone, err := query.GetStone(ctx(), stoneId)
	if err != nil {
		return err
	}
	data := stone.Data

	var jsondata datatestdata
	if err := json.Unmarshal([]byte(data), &jsondata); err != nil {
		// ignore for debugging
		return nil
	}

	patterns, err := srv.ListWithTrait()
	if err != nil {
		return err
	}
	matched := ""
	for _, pattern := range patterns {
		if isTraitsMatch(pattern.Traits, jsondata) {
			// matched
			matched = pattern.Pattern.ID
			break
		}
	}
	if matched != "" {
		err := query.UpdateStonePatternId(ctx(), dbq.UpdateStonePatternIdParams{
			PatternID: nstr(matched),
			ID:        stoneId,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

type datatestdata struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

func isTraitsMatch(traits []dbq.Trait, data datatestdata) bool {
	for _, trait := range traits {
		if trait.Path == "$.a" {
			if data.A == "" {
				return false
			}
		}
		if trait.Path == "$.b" {
			if data.B == "" {
				return false
			}
		}
		if trait.Path == "$.c" {
			if data.C == "" {
				return false
			}
		}
	}
	return true
}
